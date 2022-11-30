package domain

import "mime/multipart"

type Item struct {
	ID          int
	Name        string `binding:"required" form:"name"`
	Price       int    `binding:"required" form:"price"`
	Image       string `binding:"required" form:"image"`
	ImageName   string
	Description string `form:"description"`
	Stock       int    `form:"stock"`
	Barcode     string `binding:"required" form:"barcode"`
}

type CreateItem struct {
	ID          int
	Name        string                  `binding:"required" form:"name"`
	Price       int                     `binding:"required" form:"price"`
	Image       []*multipart.FileHeader `binding:"required" form:"image"`
	Description string                  `form:"description"`
	Stock       int                     `form:"stock"`
	Barcode     string                  `binding:"required" form:"barcode"`
}

type Pagination struct {
	Data         []Item
	Total        int64
	Perpage      int
	CurrentPage  int
	LastPage     int
	FirstPageUrl string
	NextPageUrl  string
	LastPageurl  string
	PrevPageUrl  string
	Path         string
	From         int
	To           int
}
