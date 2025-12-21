package main

import (
	contorollers "gin-fleamarket/controllers"
	"gin-fleamarket/models"
	repositoties "gin-fleamarket/repositories"
	"gin-fleamarket/services"

	"github.com/gin-gonic/gin"
)

func main() {
	items := []models.Item{
		{ID: 1, Name: "商品1", Price: 1000, Description: "説明1", SoldOut: false},
		{ID: 2, Name: "商品2", Price: 2000, Description: "説明2", SoldOut: false},
		{ID: 3, Name: "商品3", Price: 3000, Description: "説明3", SoldOut: true},
	}

	itemMemoryRepository := repositoties.NewItemMemoryRepository(items)
	itemService := services.NewItemService(itemMemoryRepository)
	itemController := contorollers.NewItemController(itemService)

	router := gin.Default()
	router.GET("/items", itemController.FindAll)
	router.GET("/items/:id", itemController.FindById)
	router.POST("/items", itemController.Create)
	router.PUT("/items/:id", itemController.Update)
	router.DELETE("/items/:id", itemController.Delete)
	router.Run() // デフォルトで0.0.0.0:8080で待機します
}
