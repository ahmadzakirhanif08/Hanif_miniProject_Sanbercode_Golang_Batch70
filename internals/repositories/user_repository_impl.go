package repositories

import (
	"context"
	"database/sql"
	"fmt"
	"miniProject/internals/models"
)


type userRepositoryImpl struct {
	DB *sql.DB
}
func NewUserRepository(db *sql.DB) UserRepository {
	return &userRepositoryImpl{DB: db}
}
func (r *userRepositoryImpl) GetByUsername(ctx context.Context, username string) (*models.User, error) {
	query := `
		SELECT 
			id, username, password, created_at, created_by, modified_at, modified_by 
		FROM users 
		WHERE username = $1
	`
	
	var u models.User
	var modifiedAt sql.NullTime
	var modifiedBy sql.NullString

	row := r.DB.QueryRowContext(ctx, query, username)
	
	err := row.Scan(
		&u.ID, &u.Username, &u.Password, &u.CreatedAt, &u.CreatedBy,
		&modifiedAt, &modifiedBy,
	)
	
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil 
		}
		return nil, fmt.Errorf("gagal scanning row GetByUsername: %w", err)
	}
	
	u.ModifiedAt = modifiedAt.Time
	u.ModifiedBy = modifiedBy.String

	return &u, nil
}

// for testing

/* func (r *userRepositoryImpl) Create(ctx context.Context, user *models.User) error {
	query := `
		INSERT INTO users (username, password, created_at, created_by) 
		VALUES ($1, $2, NOW(), $3) 
		RETURNING id
	`
	err := r.DB.QueryRowContext(ctx, query, 
		user.Username, 
		user.Password, 
		user.CreatedBy,
	).Scan(&user.ID) 
	
	if err != nil {
		return fmt.Errorf("gagal menjalankan query Create User: %w", err)
	}
	return nil
}
*/