package service

import (
	"terangasri/internal/app/domain"

	"github.com/gin-gonic/gin"
	"github.com/go-resty/resty/v2"
)

type AuthRepo interface {
	Login(auth domain.Login) (*resty.Response, interface{}, error)
	Register(auth domain.Register) (*resty.Response, interface{}, error)
}

type AuthUseCase struct {
	authRepo AuthRepo
}

func NewAuthRepo(repo AuthRepo) *AuthUseCase {
	return &AuthUseCase{repo}
}

func (au AuthUseCase) AuthLogin(ctx *gin.Context) {

	var login domain.Login
	err := ctx.ShouldBindJSON(&login)
	if err != nil {
		ctx.JSON(200, gin.H{
			"code": "404",
		})
		return
	}
	resp, result, err := au.authRepo.Login(login)

	if err != nil {
		ctx.JSON(500, gin.H{
			"code":    "500",
			"message": "Service auth off",
		})
		return
	}
	ctx.JSON(resp.StatusCode(), result)

}
func (au AuthUseCase) AuthRegister(ctx *gin.Context) {
	var register domain.Register
	err := ctx.ShouldBindJSON(&register)
	if err != nil {
		ctx.JSON(200, gin.H{
			"code": "404",
		})
		return
	}
	resp, result, err := au.authRepo.Register(register)

	if err != nil {
		ctx.JSON(500, gin.H{
			"code":    "500",
			"message": "Service auth off",
		})
		return
	}
	ctx.JSON(resp.StatusCode(), result)
}
