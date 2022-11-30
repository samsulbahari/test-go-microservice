package repository

import (
	"context"
	"item/internal/app/domain"
	"math"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type UserRepository struct {
	db *gorm.DB
}

func NewItemRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{db}
}

func (ur UserRepository) CreateItem(ctx context.Context, item domain.Item) error {
	err := ur.db.WithContext(ctx).Create(&item).Error
	return err
}

func (ur UserRepository) DeleteItem(ctx context.Context, id int) error {
	var item domain.Item
	err := ur.db.Delete(&item, id).Error
	return err
}

func (ur UserRepository) GetByid(ctx context.Context, id int) (domain.Item, error) {
	var item domain.Item
	err := ur.db.WithContext(ctx).First(&item, id).Error
	return item, err
}

func (ur UserRepository) Getdata(ctx *gin.Context, page int) ([]domain.Item, error) {

	var item []domain.Item
	offset := (page - 1) * 5
	var count int64

	ur.db.Model(&item).Count(&count)
	err := ur.db.WithContext(ctx).Limit(5).Offset(offset).Find(&item).Error

	var pagination domain.Pagination

	pagination.Data = item
	pagination.Total = count
	pagination.Perpage = 5
	if page-1 < 0 {
		//CurrentPages := fmt.Sprintf("%s/getdata?page=%d", ctx.Request.Host, 1)
		CurrentPages := 1
		pagination.CurrentPage = CurrentPages
	} else {
		CurrentPages := page - 1
		pagination.CurrentPage = CurrentPages
	}
	last_page_counts := float64(count) / float64(5)
	last_page := math.Ceil(last_page_counts)

	pagination.LastPage = int(last_page)
}
