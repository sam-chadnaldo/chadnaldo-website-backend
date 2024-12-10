package paymentrepo

import (
	"context"

	"github.com/sam-chadnaldo/chadnaldo-website-backend/internal/domain/models"
	"github.com/sam-chadnaldo/chadnaldo-website-backend/internal/repository/postgres"
)

type PaymentRepos struct {
    storage *postgres.Storage
}

func New(storage *postgres.Storage) *PaymentRepos {
    return &PaymentRepos{storage: storage}
}

// Реализация методов AuthRepository
func (r *PaymentRepos) SavePayment(ctx context.Context, user *models.User) error {
    // Реализация
	return nil
}

