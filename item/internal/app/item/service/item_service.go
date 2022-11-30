package service

import (
	"context"
	"errors"
	"fmt"
	"item/internal/app/domain"
	"item/internal/app/libraries"
	"math/rand"
	"strconv"
	"time"

	"github.com/cloudinary/cloudinary-go/api/uploader"
	"github.com/gin-gonic/gin"
)

type ItemRepository interface {
	CreateItem(ctx context.Context, item domain.Item) error
	DeleteItem(ctx context.Context, id int) error
	GetByid(ctx context.Context, id int) (domain.Item, error)
	Getdata(ctx *gin.Context, page int) ([]domain.Item, error)
}

type ItemService struct {
	ItemRepo ItemRepository
}

func NewItemService(ItemRepo ItemRepository) *ItemService {
	return &ItemService{ItemRepo}
}

func (is ItemService) CreateItemService(ctx *gin.Context, item *domain.CreateItem) (int, error) {

	rand.Seed(time.Now().UnixNano())

	// String
	charset := "abcdefghijklmnopqrstuvwxyz"

	fileName := []rune(charset)

	// Shuffling the string
	rand.Shuffle(len(fileName), func(i, j int) {
		fileName[i], fileName[j] = fileName[j], fileName[i]
	})

	// Getting random character

	file, _, _ := ctx.Request.FormFile("image")
	//fileName := charset[rand.Intn(len(charset))]
	cld, _ := libraries.Setupcloudinary()
	resp, err := cld.Upload.Upload(ctx, file, uploader.UploadParams{
		PublicID: string(fileName),
		Folder:   "goterangasri"})

	if err != nil {
		return 500, errors.New("Upload image invalid")
	}

	//item.Image = resp.SecureURL
	items := domain.Item{
		Name:        item.Name,
		Price:       item.Price,
		Image:       resp.SecureURL,
		Description: item.Description,
		Stock:       item.Stock,
		Barcode:     item.Barcode,
		ImageName:   string(fileName),
	}

	err = is.ItemRepo.CreateItem(ctx, items)
	if err != nil {
		return 500, errors.New("error insert to database")
	}

	return 200, nil

}
func (is ItemService) DeleteItem(ctx *gin.Context) (int, error) {
	id := ctx.Param("id")
	exerciseID, err := strconv.Atoi(id)
	if err != nil {
		return 200, errors.New("invalid input id")
	}

	data, err := is.ItemRepo.GetByid(ctx, exerciseID)
	if err != nil {
		return 203, errors.New("Id not found")
	}

	nameimage := fmt.Sprintf("goterangasri/%s", data.ImageName)
	cld, _ := libraries.Setupcloudinary()

	resp, err := cld.Upload.Destroy(ctx, uploader.DestroyParams{
		PublicID:     nameimage,
		ResourceType: "image"})

	if err != nil {
		return 500, errors.New("error delete image cloudinary")
	}

	fmt.Println(resp)

	err = is.ItemRepo.DeleteItem(ctx, exerciseID)
	if err != nil {
		return 500, errors.New("error delete to database")
	}

	return 200, nil
}

func (is ItemService) Getdata(ctx *gin.Context) (int, error, []domain.Item) {

	page := ctx.Request.URL.Query().Get("page")
	page_num, err := strconv.Atoi(page)

	if err != nil {
		return 500, errors.New("invalid input id"), nil
	}
	data, err := is.ItemRepo.Getdata(ctx, page_num)

	if err != nil {
		return 500, errors.New("error get data to database"), nil
	}

	return 200, nil, data

}
