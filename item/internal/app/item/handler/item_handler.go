package handler

import (
	"item/internal/app/domain"
	"item/internal/app/libraries"

	"github.com/gin-gonic/gin"
)

type ItemService interface {
	CreateItemService(*gin.Context, *domain.CreateItem) (int, error)
	DeleteItem(*gin.Context) (int, error)
	Getdata(*gin.Context) (int, error, []domain.Item)
}

type ItemHandler struct {
	ItemServ ItemService
}

func NewItemHandler(ItemServ ItemService) *ItemHandler {
	return &ItemHandler{ItemServ}
}

func (ih ItemHandler) CreateItem(ctx *gin.Context) {

	var item domain.CreateItem
	err := ctx.ShouldBind(&item)
	if err != nil {
		validation_response := libraries.Validation(err)
		ctx.JSON(200, gin.H{
			"code":    "404",
			"message": validation_response,
		})
		return
	}
	code, err := ih.ItemServ.CreateItemService(ctx, &item)
	if err != nil {
		ctx.JSON(code, gin.H{
			"code":    code,
			"message": err.Error(),
		})
		return
	}

	ctx.JSON(code, gin.H{
		"code":    code,
		"message": "Succes Create Item ",
	})
}

func (ih ItemHandler) DeleteItem(ctx *gin.Context) {
	code, err := ih.ItemServ.DeleteItem(ctx)
	if err != nil {
		ctx.JSON(200, gin.H{
			"code":    code,
			"message": err.Error(),
		})
		return
	}

	ctx.JSON(code, gin.H{
		"code":    code,
		"message": "Succes Delete Item ",
	})

}
func (ih ItemHandler) GetItem(ctx *gin.Context) {

	code, err, data := ih.ItemServ.Getdata(ctx)
	if err != nil {
		ctx.JSON(200, gin.H{
			"code":    code,
			"message": err.Error(),
		})
		return
	}
	ctx.JSON(200, gin.H{
		"code":    code,
		"message": data,
	})
	return
}
