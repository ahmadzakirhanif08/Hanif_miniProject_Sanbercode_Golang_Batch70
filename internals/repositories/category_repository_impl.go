package repositories

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"time"
	"your-github-username/quiz-bootcamp-golang/internal/models" // Ganti path sesuai mod init
)

// categoryRepositoryImpl adalah implementasi konkret dari CategoryRepository
type categoryRepositoryImpl struct {
	DB *sql.DB
}

// NewCategoryRepository membuat instance baru dari CategoryRepository
func NewCategoryRepository(db *sql.DB) CategoryRepository {
	return &categoryRepositoryImpl{DB: db}
}

func (r *categoryRepositoryImpl) GetAll(ctx context.Context) ([]models.Category, error) {
	query := "SELECT id, name, created_at, created_by, modified_at, modified_by FROM categories ORDER BY id ASC"
	
	rows, err := r.DB.QueryContext(ctx, query)
	if err != nil {
		return nil, fmt.Errorf("gagal menjalankan query GetAll: %w", err)
	}
	defer rows.Close()

	var categories []models.Category
	for rows.Next() {
		var c models.Category
		var modifiedAt sql.NullTime
		var modifiedBy sql.NullString
		
		err := rows.Scan(
			&c.ID, &c.Name, &c.CreatedAt, &c.CreatedBy, 
			&modifiedAt, &modifiedBy,
		)
		if err != nil {
			return nil, fmt.Errorf("gagal scanning row: %w", err)
		}
		
		// Mengatasi kolom yang nullable (modified_at, modified_by)
		c.ModifiedAt = modifiedAt.Time
		c.ModifiedBy = modifiedBy.String
		
		categories = append(categories, c)
	}

	if err = rows.Err(); err != nil {
		return nil, fmt.Errorf("error setelah iterasi rows: %w", err)
	}

	return categories, nil
}

// GetByID mengembalikan kategori berdasarkan ID
func (r *categoryRepositoryImpl) GetByID(ctx context.Context, id int) (*models.Category, error) {
	query := "SELECT id, name, created_at, created_by, modified_at, modified_by FROM categories WHERE id = $1"
	
	var c models.Category
	var modifiedAt sql.NullTime
	var modifiedBy sql.NullString

	row := r.DB.QueryRowContext(ctx, query, id)
	
	err := row.Scan(
		&c.ID, &c.Name, &c.CreatedAt, &c.CreatedBy,
		&modifiedAt, &modifiedBy,
	)
	
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			// Mengembalikan nil (tanpa error) jika data tidak ditemukan, 
			// sehingga service layer bisa memberikan pesan "tidak ditemukan" yang tepat.
			return nil, nil 
		}
		return nil, fmt.Errorf("gagal scanning row GetByID: %w", err)
	}
	
	// Mengatasi kolom yang nullable
	c.ModifiedAt = modifiedAt.Time
	c.ModifiedBy = modifiedBy.String

	return &c, nil
}

// Create menambahkan kategori baru
func (r *categoryRepositoryImpl) Create(ctx context.Context, category *models.Category) error {
	query := `
		INSERT INTO categories (name, created_at, created_by) 
		VALUES ($1, $2, $3) 
		RETURNING id
	`
	// Isi kolom yang dikelola di layer service/repo
	category.CreatedAt = time.Now() 
	
	err := r.DB.QueryRowContext(ctx, query, 
		category.Name, 
		category.CreatedAt,
		category.CreatedBy,
	).Scan(&category.ID) // Mengambil ID yang baru di-generate
	
	if err != nil {
		return fmt.Errorf("gagal menjalankan query Create: %w", err)
	}
	return nil
}

func (r *categoryRepositoryImpl) Delete(ctx context.Context, id int) error {
	query := "DELETE FROM categories WHERE id = $1"
	
	result, err := r.DB.ExecContext(ctx, query, id)
	if err != nil {
		return fmt.Errorf("gagal menjalankan query Delete: %w", err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("gagal mendapatkan RowsAffected: %w", err)
	}

	if rowsAffected == 0 {
		return errors.New("gagal menghapus: kategori tidak ditemukan")
	}

	return nil
}


func (r *categoryRepositoryImpl) Update(ctx context.Context, category *models.Category) error {
	query := `
		UPDATE categories 
		SET name = $1, modified_at = $2, modified_by = $3
		WHERE id = $4
	`
	// Isi kolom modified_at di layer service/repo
	category.ModifiedAt = time.Now()
	
	result, err := r.DB.ExecContext(ctx, query, 
		category.Name, 
		category.ModifiedAt, 
		category.ModifiedBy, 
		category.ID,
	)

	if err != nil {
		return fmt.Errorf("gagal menjalankan query Update: %w", err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("gagal mendapatkan RowsAffected (Update): %w", err)
	}

	if rowsAffected == 0 {
		return errors.New("gagal update: kategori tidak ditemukan")
	}
	
	return nil
}