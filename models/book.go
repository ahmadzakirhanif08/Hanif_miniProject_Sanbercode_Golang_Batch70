package models

import (
	"time"
)

type Book struct {
	ID          int       `json:"id"`
	Title       string    `json:"title" binding:"required"`
	CategoryID  int       `json:"category_id" binding:"required"`
	Description string    `json:"description"`
	ImageURL    string    `json:"image_url"`
	ReleaseYear int       `json:"release_year" binding:"required,gte=1980,lte=2024"` 
	Price       int       `json:"price"`
	TotalPage   int       `json:"total_page" binding:"required,gt=0"`
	Thickness   string    `json:"thickness"`
	CreatedAt   time.Time `json:"created_at"` 
	CreatedBy   string    `json:"created_by"`
	ModifiedAt  time.Time `json:"modified_at"`
	ModifiedBy  string    `json:"modified_by"`
}