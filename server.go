package main

import (
	"fmt"
	"net/http"
	"shopaholic-service/controller"
	"shopaholic-service/types"
	"shopaholic-service/utilities"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.Use(utilities.Logger())

	// Recovery middleware recovers from any panics and writes a 500 if there was one.
	router.Use(gin.Recovery())

	router.LoadHTMLGlob("templates/*")
	// Health check
	router.GET("/ping", func(c *gin.Context) {
		c.String(http.StatusOK, "Hello world!")
	})

	// Simple group: v1
	v1 := router.Group("/v1")
	{
		// Create a  new shopping list
		v1.POST("/shopping-lists", func(c *gin.Context) {
			var list types.ShoppingList
			c.ShouldBind(&list)
			res, _ := controller.CreateShoppingList(list)
			c.JSON(http.StatusOK, res)
		})

		// Get all shopping lists
		v1.GET("/shopping-lists", func(c *gin.Context) {
			lists, err := controller.GetShoppingLists()
			if err != nil {
				fmt.Println(err)
				return
			}
			c.JSON(http.StatusOK, lists)
		})

		// Get a specific shopping list
		v1.GET("/shopping-lists/:listID", func(c *gin.Context) {
			listID := c.Param("listID")
			list, _ := controller.GetShoppingList(listID)
			c.JSON(http.StatusOK, list)
		})

		// Remove an item to a shopping list
		v1.DELETE("/shopping-lists/:listID", func(c *gin.Context) {
			listID := c.Param("listID")
			res, _ := controller.RemoveShoppingList(listID)
			c.JSON(http.StatusOK, res)
		})

		// Add an item to a shopping list
		v1.POST("/shopping-lists/:listID/items", func(c *gin.Context) {
			listID := c.Param("listID")
			var item types.ListItem
			c.ShouldBind(&item)
			addedItem, err := controller.AddItemToShoppingList(listID, item)
			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"errors": err.Error()})
				return
			}
			c.JSON(http.StatusOK, addedItem)
		})

		// Update an item in a shopping list
		v1.PATCH("/shopping-lists/:listID/items/:itemID", func(c *gin.Context) {
			listID := c.Param("listID")
			itemID := c.Param("itemID")
			var item types.ListItem
			c.ShouldBind(&item)
			res, _ := controller.UpdateShoppingListItem(listID, itemID, item)
			c.JSON(http.StatusOK, res)
		})

		// Remove an item from a shopping list
		v1.DELETE("/shopping-lists/:listID/items/:itemID", func(c *gin.Context) {
			listID := c.Param("listID")
			itemID := c.Param("itemID")
			res, _ := controller.RemoveItemFromShoppingList(listID, itemID)
			c.JSON(http.StatusOK, res)
		})

	}

	// Remove an item from a shopping list
	router.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusNotFound, gin.H{"code": "PAGE_NOT_FOUND", "message": "Page not found"})
	})
	router.Run(":8080")
}
