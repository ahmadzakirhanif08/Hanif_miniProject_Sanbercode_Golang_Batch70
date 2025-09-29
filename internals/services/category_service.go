package services

import (
	"context"
	"errors"
	"miniProject/internals/models"
	"miniProject/internals/repositories"
)

// CategoryService mendefinisikan business logic untuk entitas Category
type CategoryService interface {
	FindAll(ctx context.Context) ([]models.Category, error)
	FindByID(ctx context.Context, id int) (*models.Category, error)
	Create(ctx context.Context, name string, createdBy string) (*models.Category, error)
	Delete(ctx context.Context, id int) error
	Update(ctx context.Context, id int, name string, modifiedBy string) error
}

// categoryServiceImpl adalah implementasi konkret dari CategoryService
type categoryServiceImpl struct {
	repo repositories.CategoryRepository
}

// NewCategoryService membuat instance baru dari CategoryService
func NewCategoryService(repo repositories.CategoryRepository) CategoryService {
	return &categoryServiceImpl{repo: repo}
}

// -----------------------------------------------------------------------------
// Implementasi Methods
// -----------------------------------------------------------------------------

func (s *categoryServiceImpl) FindAll(ctx context.Context) ([]models.Category, error) {
	// Logic bisnis minimal: langsung panggil repository
	return s.repo.GetAll(ctx)
}

func (s *categoryServiceImpl) FindByID(ctx context.Context, id int) (*models.Category, error) {
	// Logic bisnis minimal: langsung panggil repository
	return s.repo.GetByID(ctx, id)
}

func (s *categoryServiceImpl) Create(ctx context.Context, name string, createdBy string) (*models.Category, error) {
	// 1. Lakukan validasi tambahan jika ada (misal: cek apakah nama sudah ada)
	// Untuk saat ini, kita asumsikan validasi 'required' sudah dilakukan di handler/binding.
	
	// 2. Siapkan Model untuk disimpan
	newCategory := models.Category{
		Name:      name,
		CreatedBy: createdBy,
	}

	// 3. Panggil Repository
	err := s.repo.Create(ctx, &newCategory)
	if err != nil {
		return nil, err
	}

	return &newCategory, nil
}

func (s *categoryServiceImpl) Delete(ctx context.Context, id int) error {
	// Persyaratan: Berikan message error jika menghapus data kategori yang tidak tersedia [cite: 147]
	
	// 1. Cek keberadaan kategori sebelum menghapus
	existingCategory, err := s.repo.GetByID(ctx, id)
	if err != nil {
		// Asumsikan repository mengembalikan error spesifik (misal: sql.ErrNoRows) jika tidak ditemukan
		return errors.New("data kategori tidak tersedia atau terjadi error database")
	}
	if existingCategory == nil {
		return errors.New("gagal menghapus: data kategori tidak ditemukan")
	}

	// 2. Panggil Repository untuk menghapus
	return s.repo.Delete(ctx, id)
}

func (s *categoryServiceImpl) Update(ctx context.Context, id int, name string, modifiedBy string) error {
	// Persyaratan: Berikan message error jika mengupdate data kategori yang tidak tersedia [cite: 148]

	// 1. Cek keberadaan kategori sebelum memperbarui
	existingCategory, err := s.repo.GetByID(ctx, id)
	if err != nil {
		return errors.New("data kategori tidak tersedia atau terjadi error database")
	}
	if existingCategory == nil {
		return errors.New("gagal update: data kategori tidak ditemukan")
	}

	// 2. Update field
	existingCategory.Name = name
	existingCategory.ModifiedBy = modifiedBy
	
	// ModifiedAt akan diisi di repository atau di level database (default NOW())
	
	// 3. Panggil Repository untuk memperbarui
	return s.repo.Update(ctx, existingCategory)
}