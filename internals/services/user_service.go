package services

import (
	"context"
	"errors"
	"miniProject/internals/models"
	"miniProject/internals/repositories"
	"golang.org/x/crypto/bcrypt"
)

type UserService interface {
	Login(ctx context.Context, username string, password string) (string, error)
	// ...
}

type userServiceImpl struct {
	repo repositories.UserRepository
}

func NewUserService(repo repositories.UserRepository) UserService {
	return &userServiceImpl{repo: repo}
}

// Login menangani business logic untuk login pengguna dan menghasilkan JWT
func (s *userServiceImpl) Login(ctx context.Context, username string, password string) (string, error) {
	// 1. Dapatkan pengguna dari database
	user, err := s.repo.GetByUsername(ctx, username)
	if err != nil {
		return "", errors.New("gagal saat mencari pengguna")
	}
	if user == nil {
		return "", errors.New("username atau password salah")
	}

	// 2. Verifikasi Password (asumsi password di DB adalah hash bcrypt)
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		if errors.Is(err, bcrypt.ErrMismatchedHashAndPassword) {
			return "", errors.New("username atau password salah")
		}
		return "", errors.New("gagal memverifikasi password")
	}

	// 3. Generate JWT
	tokenString, err := utils.GenerateToken(user.Username)
	if err != nil {
		return "", errors.New("gagal menghasilkan token")
	}

	return tokenString, nil
}