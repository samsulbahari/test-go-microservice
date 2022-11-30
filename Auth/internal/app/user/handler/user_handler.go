package handler

import (
	"auth/internal/app/domain"
	"auth/internal/app/libraries"
	"context"

	"github.com/gin-gonic/gin"
)

type UserService interface {
	Login(context.Context, domain.Login) (string, error)
	Register(context.Context, domain.User) (int, error)
}

type UserHandler struct {
	userService UserService
}

func NewUserHandler(userService UserService) *UserHandler {
	return &UserHandler{userService}
}

func (uh UserHandler) Login(ctx *gin.Context) {
	var param domain.Login
	err := ctx.ShouldBindJSON(&param)
	if err != nil {
		validation_response := libraries.Validation(err)
		ctx.JSON(200, gin.H{
			"code":    "404",
			"message": validation_response,
		})
		return
	}

	token, err := uh.userService.Login(ctx, param)
	if err != nil {
		ctx.JSON(200, gin.H{
			"code":    403,
			"message": err.Error(),
		})
		return
	}
	ctx.JSON(200, gin.H{
		"code":    200,
		"message": "Success Login",
		"token":   token,
	})

}
func (uh UserHandler) Register(ctx *gin.Context) {
	var user domain.User
	err := ctx.ShouldBindJSON(&user)
	if err != nil {
		validation_response := libraries.Validation(err)
		ctx.JSON(200, gin.H{
			"code":    "404",
			"message": validation_response,
		})
		return
	}

	code, err := uh.userService.Register(ctx, user)
	if err != nil {
		ctx.JSON(200, gin.H{
			"code":    code,
			"message": err.Error(),
		})
		return
	}

	ctx.JSON(200, gin.H{
		"code":    code,
		"message": "Success Create user",
	})
	return

}
