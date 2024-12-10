package authrepo

import (
	"context"

	"github.com/sam-chadnaldo/chadnaldo-website-backend/internal/domain/models"
	"github.com/sam-chadnaldo/chadnaldo-website-backend/internal/repository/postgres"
)

type AuthRepository struct {
    storage *postgres.Storage
}

func NewUserRepository(storage *postgres.Storage) *AuthRepository {
    return &AuthRepository{storage: storage}
}

// Реализация методов AuthRepository
func (r *AuthRepository) Register(ctx context.Context, user *models.User) error {
    // Реализация
	return nil
}

func (r *AuthRepository) Login(ctx context.Context, email, password string) (*models.User, error) {
    // Реализация
	return &models.User{}, nil
}