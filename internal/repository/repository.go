package repository

import (
	"github.com/sam-chadnaldo/chadnaldo-website-backend/internal/repository/postgres"
	apprepo "github.com/sam-chadnaldo/chadnaldo-website-backend/internal/repository/postgres/app"
	paymentrepo "github.com/sam-chadnaldo/chadnaldo-website-backend/internal/repository/postgres/payment"
)

type AppRepository interface{
	// App(ctx context.Context, id int) (models.App, error)
}

type PaymentRepository interface {

}
// Добавьте другие интерфейсы по необходимости

type Repository struct {
	AppRepository
	PaymentRepository
}

func NewRepository(storage *postgres.Storage) *Repository {
    return &Repository{
		AppRepository: apprepo.NewAppRepository(storage),
		PaymentRepository: paymentrepo.New(storage),
    }
}
