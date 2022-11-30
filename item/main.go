package main

import (
	"item/internal/app/database"
	"item/internal/app/item/handler"
	"item/internal/app/item/repository"
	"item/internal/app/item/service"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("error loading env")
	}
	router := gin.Default()
	db := database.ConnectDatabase()

	ItemRepo := repository.NewItemRepository(db)
	ItemService := service.NewItemService(ItemRepo)
	ItemHandler := handler.NewItemHandler(ItemService)

	router.POST("/create", ItemHandler.CreateItem)
	router.POST("/delete/:id", ItemHandler.DeleteItem)
	router.GET("/getdata", ItemHandler.GetItem)
	router.Run()
}
