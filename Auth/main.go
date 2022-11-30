package main

import (
	"log"

	"auth/internal/app/database"
	"auth/internal/app/user/handler"
	"auth/internal/app/user/repository"
	"auth/internal/app/user/service"

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
	userRepo := repository.NewUserRepository(db)
	userService := service.NewUserService(userRepo)
	UserHandler := handler.NewUserHandler(userService)

	router.POST("/login", UserHandler.Login)
	router.POST("/register", UserHandler.Register)

	router.Run()

}
