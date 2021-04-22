package main

import (
	"fmt"
	"net/http"
	"shopaholic-service/data"
	"shopaholic-service/types"
	"shopaholic-service/utilities"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.Use(utilities.Logger())
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
	router.GET("/shopping-lists/:listID", func(c *gin.Context) {
		listID := c.Param("listID")
		list, _ := data.GetShoppingList(listID)
		c.JSON(http.StatusOK, list)
	})

	// Remove an item to a shopping list
	router.DELETE("/shopping-lists/:listID", func(c *gin.Context) {
		listID := c.Param("listID")
		res, _ := data.RemoveShoppingList(listID)
		c.JSON(http.StatusOK, res)
	})

	// Add an item to a shopping list
	router.POST("/shopping-lists/:listID/items", func(c *gin.Context) {
		listID := c.Param("listID")
		var item types.ListItem
		c.ShouldBind(&item)
		addedItem, err := data.AddItemToShoppingList(listID, item)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"errors": err.Error()})
			return
		}
		c.JSON(http.StatusOK, addedItem)
	})

	// Update an item in a shopping list
	router.PATCH("/shopping-lists/:listID/items/:itemID", func(c *gin.Context) {
		listID := c.Param("listID")
		itemID := c.Param("itemID")
		var item types.ListItem
		c.ShouldBind(&item)
		res, _ := data.UpdateShoppingListItem(listID, itemID, item)
		c.JSON(http.StatusOK, res)
	})

	// Remove an item from a shopping list
	router.DELETE("/shopping-lists/:listID/items/:itemID", func(c *gin.Context) {
		listID := c.Param("listID")
		itemID := c.Param("itemID")
		res, _ := data.RemoveItemFromShoppingList(listID, itemID)
		c.JSON(http.StatusOK, res)
	})

	router.Run(":8080")
}
