package repositories

import (
	"context"
	"miniProject/internals/models"
)

type UserRepository interface {
	GetByUsername(ctx context.Context, username string) (*models.User, error)
}