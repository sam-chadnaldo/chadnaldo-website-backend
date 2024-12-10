package services

import (
	"context"
	"log/slog"

	"github.com/sam-chadnaldo/chadnaldo-website-backend/internal/repository"
	"github.com/sam-chadnaldo/chadnaldo-website-backend/internal/services/payment"
)

type PaymetnService interface{
	GetPaymentInvoiceLink(ctx context.Context, userID string, tokensAmount int64) (string, error)
}

type Services struct {
	payment PaymetnService
}

func NewServices(repo *repository.PaymentRepository, logger *slog.Logger) *Services{
	return&Services{
		payment: payment.New(repo, logger),
	}
}