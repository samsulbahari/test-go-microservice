package main

import (
	"terangasri/internal/app/authService/repository"
	"terangasri/internal/app/authService/service"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	userRepo := repository.NewUserServiceRepository()
	userUserCase := service.NewAuthRepo(userRepo)
	r.POST("login", userUserCase.AuthLogin)
	r.POST("register", userUserCase.AuthRegister)
	r.Run(":1234")
}
