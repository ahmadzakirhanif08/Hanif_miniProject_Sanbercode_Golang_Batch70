package controllers

import (
	"miniProject/config" 
	"miniProject/models"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

func determineThickness(totalPage int) string {
	if totalPage > 100 {
		return "tebal"
	}

	return "tipis"
}


func CreateBook(c *gin.Context) {
	var newBook models.Book
	if err := c.ShouldBindJSON(&newBook); err != nil {
		// Validasi bawaan Gin (termasuk release_year) 
		c.JSON(http.StatusBadRequest, gin.H{"error": "Validation failed", "details": err.Error()})
		return
	}

	newBook.Thickness = determineThickness(newBook.TotalPage)


	username, _ := c.Get(gin.AuthUserKey) 
	newBook.CreatedAt = time.Now()
	newBook.CreatedBy = username.(string)


	query := `INSERT INTO books (title, category_id, description, image_url, release_year, price, total_page, thickness, created_at, created_by) 
	          VALUES RETURNING id`
	
	err := config.DB.QueryRow(query, 
		newBook.Title, newBook.CategoryID, newBook.Description, newBook.ImageURL, newBook.ReleaseYear, 
		newBook.Price, newBook.TotalPage, newBook.Thickness, newBook.CreatedAt, newBook.CreatedBy).Scan(&newBook.ID)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create book"})
		return
	}

	c.JSON(http.StatusCreated, newBook)
}


func FindBooks(c *gin.Context) {
	var books []models.Book
	
	rows, err := config.DB.Query("SELECT id, title, release_year, thickness, total_page FROM books ORDER BY id ASC")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve books"})
		return
	}
	defer rows.Close()

	for rows.Next() {
		var book models.Book
		if err := rows.Scan(&book.ID, &book.Title, &book.ReleaseYear, &book.Thickness, &book.TotalPage); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Error scanning book data"})
			return
		}
		books = append(books, book)
	}

	c.JSON(http.StatusOK, books)
}


func DeleteBook(c *gin.Context) {
	idStr := c.Param("id")
	id, _ := strconv.Atoi(idStr)
	

	var exists bool
	checkQuery := `SELECT EXISTS(SELECT 1 FROM books WHERE id = $1)`
	config.DB.QueryRow(checkQuery, id).Scan(&exists)
	
	if !exists {
		c.JSON(http.StatusNotFound, gin.H{"error": "Message error: Book not available for deletion."})
		return
	}
	
	// Lakukan DELETE
	result, err := config.DB.Exec("DELETE FROM books WHERE id = $1", id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete book"})
		return
	}

	rowsAffected, _ := result.RowsAffected()
	if rowsAffected == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "Message error: Book not available for deletion."})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Book deleted successfully"})
}


func GetBookByID(c *gin.Context) {
	idStr := c.Param("id")
	id, _ := strconv.Atoi(idStr)
	
	var book models.Book
	query := `SELECT id, title, description, release_year, price, total_page, thickness FROM books WHERE id = $1`
	err := config.DB.QueryRow(query, id).Scan(&book.ID, &book.Title, &book.Description, &book.ReleaseYear, &book.Price, &book.TotalPage, &book.Thickness)

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Book not found"})
		return
	}

	c.JSON(http.StatusOK, book)
}