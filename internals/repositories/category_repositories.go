package repositories

import (
	"context"
	"miniProject/internals/models"
)

// CategoryRepository mendefinisikan operasi CRUD untuk model Category
type CategoryRepository interface {
	// GetAll mengembalikan semua kategori
	GetAll(ctx context.Context) ([]models.Category, error)
	
	// GetByID mengembalikan kategori berdasarkan ID
	GetByID(ctx context.Context, id int) (*models.Category, error)
	
	// Create menambahkan kategori baru ke database
	Create(ctx context.Context, category *models.Category) error
	
	// Delete menghapus kategori berdasarkan ID
	Delete(ctx context.Context, id int) error
	
	// Update memperbarui kategori yang sudah ada
	Update(ctx context.Context, category *models.Category) error
}