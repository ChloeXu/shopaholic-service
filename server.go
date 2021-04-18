package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"shopaholic-service/data"
	"shopaholic-service/types"
	"strconv"

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
		// // load thru html content
		// c.HTML(http.StatusOK, "tech_design.tmpl", gin.H{
		// 	"title": designPage.Title,
		// 	"textBody":  designPage.Text,
		// })
		c.String(http.StatusOK, designPage.Text)
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
		lists, err := data.GetShoppingLists()
		if err != nil {
			fmt.Println(err)
			return
		}
		for i := 0; i < len(lists); i++ {
			listIdInt, _ := strconv.Atoi(listId)
			if lists[i].ID == listIdInt {

				c.JSON(http.StatusOK, lists[i])
				return
			}
		}
		c.JSON(http.StatusNotFound, gin.H{"status": http.StatusNotFound})
	})

	// However, this one will match /user/john/ and also /user/john/send
	// If no other routers match /user/john, it will redirect to /user/john/
	router.GET("/user/:name/*action", func(c *gin.Context) {
		name := c.Param("name")
		action := c.Param("action")
		message := name + " is " + action
		c.String(http.StatusOK, message)
	})

	router.Run(":8080")
}
