package apprepo

import (
	"github.com/sam-chadnaldo/chadnaldo-website-backend/internal/repository/postgres"
)

type AppRepository struct {
    storage *postgres.Storage
}

func NewAppRepository(storage *postgres.Storage) *AppRepository {
    return &AppRepository{storage: storage}
}

// func (s *AppRepository) App(ctx context.Context, id int) (models.App, error) {
// 	const op = "repository.postgres.app.App"

// 	stmt, err := s.storage.DB.Prepare("SELECT id, name, secret FROM apps WHERE id = ?")
// 	if err != nil {
// 		return models.App{}, fmt.Errorf("%s: %w", op, err)
// 	}

// 	row := stmt.QueryRowContext(ctx, id)

// 	var app models.App
// 	err = row.Scan(&app.ID, &app.Name, &app.Secret)
// 	if err != nil {
// 		if errors.Is(err, sql.ErrNoRows) {
// 			return models.App{}, fmt.Errorf("%s: %w", op, postgres.ErrAppNotFound)
// 		}

// 		return models.App{}, fmt.Errorf("%s: %w", op, err)
// 	}

// 	return app, nil
// }