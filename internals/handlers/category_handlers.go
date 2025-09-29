package handlers

import (
	"log"
	"net/http"
	"strconv"

	"miniProject/internals/services"
	
	"github.com/gin-gonic/gin"
)

type CategoryHandler struct {
	categoryService services.CategoryService
}

func NewCategoryHandler(s services.CategoryService) *CategoryHandler {
	return &CategoryHandler{categoryService: s}
}

const DefaultUser = "system_user"

// GetAll menampilkan seluruh kategori
func (h *CategoryHandler) GetAll(c *gin.Context) {
	categories, err := h.categoryService.FindAll(c.Request.Context())
	if err != nil {
		// Log error server
		log.Printf("Error FindAll categories: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve categories"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Success retrieving all categories",
		"data": categories,
	})
}

// GetByID menampilkan detail kategori
func (h *CategoryHandler) GetByID(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid category ID format"})
		return
	}

	category, err := h.categoryService.FindByID(c.Request.Context(), id)
	if err != nil {
		log.Printf("Error FindByID category: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve category detail"})
		return
	}

	if category == nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Category not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Success retrieving category detail",
		"data": category,
	})
}

// Create menambahkan kategori baru
func (h *CategoryHandler) Create(c *gin.Context) {
	var input struct {
		Name string `json:"name" binding:"required"`
	}

	// 1. Binding dan Validasi Input
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// 2. Panggil Service Layer
	// Menggunakan DefaultUser sebagai placeholder untuk created_by
	newCategory, err := h.categoryService.Create(c.Request.Context(), input.Name, DefaultUser) 
	if err != nil {
		log.Printf("Error Create category: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create category"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "Category successfully created",
		"data": newCategory,
	})
}

// Delete menghapus kategori
func (h *CategoryHandler) Delete(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid category ID format"})
		return
	}

	// Panggil Service Layer
	err = h.categoryService.Delete(c.Request.Context(), id)
	if err != nil {
		// Persyaratan: Berikan message error jika menghapus data kategori yang tidak tersedia.
		// Kami asumsikan service mengembalikan error spesifik/informatif.
		if err.Error() == "gagal menghapus: data kategori tidak ditemukan" {
			c.JSON(http.StatusNotFound, gin.H{"error": "Gagal menghapus: data kategori tidak tersedia."})
			return
		}

		log.Printf("Error Delete category: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete category"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Category successfully deleted"})
}

// internal/handlers/category_handler.go

// Update memperbarui kategori yang sudah ada
func (h *CategoryHandler) Update(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid category ID format"})
		return
	}

	var input struct {
		Name string `json:"name" binding:"required"`
	}

	// 1. Binding dan Validasi Input
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Input tidak valid: Name wajib diisi"})
		return
	}

	// 2. Ambil username dari Context (yang disetel oleh JWT Middleware)
	modifiedBy, exists := c.Get("username")
	if !exists {
		// Logika fallback jika middleware gagal setel username
		modifiedBy = DefaultUser // Placeholder
	}

	// 3. Panggil Service Layer
	err = h.categoryService.Update(c.Request.Context(), id, input.Name, modifiedBy.(string))
	if err != nil {
		// Persyaratan: Berikan message error jika mengupdate data kategori yang tidak tersedia.
		if err.Error() == "gagal update: data kategori tidak ditemukan" {
			c.JSON(http.StatusNotFound, gin.H{"error": "Gagal update: data kategori tidak tersedia."})
			return
		}

		log.Printf("Error Update category: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update category"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Category successfully updated"})
}