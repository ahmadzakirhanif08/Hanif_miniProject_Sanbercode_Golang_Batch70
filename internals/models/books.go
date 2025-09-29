package models

import "time"


type Book struct {
	ID          int       `json:"id"`
	Title       string    `json:"title" binding:"required"`
	CategoryID  int       `json:"category_id" binding:"required"`
	Description string    `json:"description"`
	ImageURL    string    `json:"image_url"`
	ReleaseYear int       `json:"release_year" binding:"required,min=1980,max=2024"`
	Price       int       `json:"price" binding:"required,min=0"`
	TotalPage   int       `json:"total_page" binding:"required,min=1"`
	Thickness   string    `json:"thickness"`
	CreatedAt   time.Time `json:"created_at"`
	CreatedBy   string    `json:"created_by"`
	ModifiedAt  time.Time `json:"modified_at"`
	ModifiedBy  string    `json:"modified_by"`
}

type BookCreateRequest struct {
	Title       string `json:"title" binding:"required"`
	CategoryID  int    `json:"category_id" binding:"required"`
	Description string `json:"description"`
	ImageURL    string `json:"image_url"`
	ReleaseYear int `json:"release_year" binding:"required,min=1980,max=2024"`
	Price       int `json:"price" binding:"required,min=0"`
	TotalPage   int `json:"total_page" binding:"required,min=1"`
}