package controllers

import (
	"miniProject/config" 
	"miniProject/models"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

func CreateCategory(c *gin.Context) {
	var newCategory models.Category
	if err := c.ShouldBindJSON(&newCategory); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	username, _ := c.Get(gin.AuthUserKey) 
	
	newCategory.CreatedAt = time.Now()
	newCategory.CreatedBy = username.(string)

	query := `INSERT INTO categories (name, created_at, created_by) VALUES ($1, $2, $3) RETURNING id`
	err := config.DB.QueryRow(query, newCategory.Name, newCategory.CreatedAt, newCategory.CreatedBy).Scan(&newCategory.ID)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create category"})
		return
	}

	c.JSON(http.StatusCreated, newCategory)
}

func FindCategories(c *gin.Context) {
	var categories []models.Category
	
	rows, err := config.DB.Query("SELECT id, name, created_at, created_by FROM categories ORDER BY id ASC")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve categories"})
		return
	}
	defer rows.Close()

	for rows.Next() {
		var cat models.Category
		if err := rows.Scan(&cat.ID, &cat.Name, &cat.CreatedAt, &cat.CreatedBy); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Error scanning category data"})
			return
		}
		categories = append(categories, cat)
	}

	c.JSON(http.StatusOK, categories)
}

func GetCategoryByID(c *gin.Context) {
	idStr := c.Param("id")
	id, _ := strconv.Atoi(idStr)
	
	var category models.Category
	query := `SELECT id, name, created_at, created_by FROM categories WHERE id = $1`
	err := config.DB.QueryRow(query, id).Scan(&category.ID, &category.Name, &category.CreatedAt, &category.CreatedBy)

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Category not found"}) 
		return
	}

	c.JSON(http.StatusOK, category)
}

func DeleteCategory(c *gin.Context) {
	idStr := c.Param("id")
	id, _ := strconv.Atoi(idStr)
	
	var exists bool
	checkQuery := `SELECT EXISTS(SELECT 1 FROM categories WHERE id = $1)`
	config.DB.QueryRow(checkQuery, id).Scan(&exists)
	
	if !exists {
		c.JSON(http.StatusNotFound, gin.H{"error": "Message error: Category not available for deletion."})
		return
	}
	
	result, err := config.DB.Exec("DELETE FROM categories WHERE id = $1", id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete category"})
		return
	}

	rowsAffected, _ := result.RowsAffected()
	if rowsAffected == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "Message error: Category not available for deletion."})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Category deleted successfully"})
}

func FindBooksByCategory(c *gin.Context) {
	categoryIDStr := c.Param("id")
	categoryID, _ := strconv.Atoi(categoryIDStr)
	
	var books []models.Book

	query := `SELECT id, title, description, release_year, thickness FROM books WHERE category_id = $1`
	rows, err := config.DB.Query(query, categoryID)
	
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve books"})
		return
	}
	defer rows.Close()

	for rows.Next() {
		var book models.Book
		if err := rows.Scan(&book.ID, &book.Title, &book.Description, &book.ReleaseYear, &book.Thickness); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Error scanning book data"})
			return
		}
		books = append(books, book)
	}

	c.JSON(http.StatusOK, books)
}