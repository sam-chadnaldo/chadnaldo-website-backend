package payment

import (
	"context"
	"log/slog"
	"time"

	"github.com/sam-chadnaldo/chadnaldo-website-backend/internal/repository"
)

type PaymentService struct {
	repo *repository.PaymentRepository
	log *slog.Logger
	tokenTTL    time.Duration
}

func New(repo *repository.PaymentRepository, logger *slog.Logger) *PaymentService{


	return &PaymentService{
		repo: repo,
		log: logger,
	}
}

func (s *PaymentService) GetPaymentInvoiceLink(ctx context.Context, userID string, tokensAmount int64) (string, error){
	const op = "services.payment.GetPaymentInvoiceLink"

	logger := s.log.With( "op", op )
	logger.Info("Neew to implement")
	return "", nil
}