package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"shopaholic-service/data"
	"shopaholic-service/types"

	"github.com/gin-gonic/gin"
)

func testingHandler(w http.ResponseWriter, r *http.Request) {
	sampleList := types.ShoppingList{
		ID:        1,
		Name:      "HMart",
		CreatedAt: "2021-04-19",
		Items:     []types.ListItem{},
	}
	b, err := json.Marshal(sampleList)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Fprint(w, string(b))
}

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hi there, I love %s!", r.URL.Path[1:])
}

func main() {
	router := gin.Default()

	router.LoadHTMLGlob("templates/*")
	// Health check
	router.GET("/ping", func(c *gin.Context) {
		c.String(http.StatusOK, "Hello world!")
	})

	// Get Tech design
	router.GET("/tech-design", func(c *gin.Context) {
		designPage, _ := loadTechDesign()
		c.String(http.StatusOK, designPage.Text)
	})

	// Create a  new shopping list
	router.POST("/shopping-lists", func(c *gin.Context) {
		var list types.ShoppingList
		c.ShouldBind(&list)
		res, _ := data.CreateShoppingList(list)
		c.JSON(http.StatusOK, res)
	})

	// Get all shopping lists
	router.GET("/shopping-lists", func(c *gin.Context) {
		lists, err := data.GetShoppingLists()
		if err != nil {
			fmt.Println(err)
			return
		}
		c.JSON(http.StatusOK, lists)
	})

	// Get a specific shopping list
	router.GET("/shopping-lists/:listId", func(c *gin.Context) {
		listId := c.Param("listId")
		list, _ := data.GetShoppingList(listId)
		c.JSON(http.StatusOK, list)
	})

	// Remove an item to a shopping list
	router.DELETE("/shopping-lists/:listId", func(c *gin.Context) {
		listId := c.Param("listId")
		res, _ := data.RemoveShoppingList(listId)
		c.JSON(http.StatusOK, res)
	})

	// Add an item to a shopping list
	router.POST("/shopping-lists/:listId/items", func(c *gin.Context) {
		listId := c.Param("listId")
		var item types.ListItem
		c.ShouldBind(&item)
		addedItem, _ := data.AddItemToShoppingList(listId, item)
		c.JSON(http.StatusOK, addedItem)
	})

	// Remove an item to a shopping list
	router.PATCH("/shopping-lists/:listId/items/:itemId", func(c *gin.Context) {
		listId := c.Param("listId")
		itemId := c.Param("itemId")
		var item types.ListItem
		c.ShouldBind(&item)
		res, _ := data.UpdateShoppingListItem(listId, itemId, item)
		c.JSON(http.StatusOK, res)
	})

	// Remove an item to a shopping list
	router.DELETE("/shopping-lists/:listId/items/:itemId", func(c *gin.Context) {
		listId := c.Param("listId")
		itemId := c.Param("itemId")
		res, _ := data.RemoveItemFromShoppingList(listId, itemId)
		c.JSON(http.StatusOK, res)
	})

	router.Run(":8080")
}
