package postgres

import (
	"fmt"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq" // Явно импортируем драйвер PostgreSQL
)

type Config struct {
	URL      string
	Host     string
	Port     string
	Username string
	Password string
	DBName   string
	SSLMode  string
}

type Storage struct {
	DB *sqlx.DB
}

func New(cfg *Config) (*Storage, error) {
	const op = "storage.postgres.postgres.NewDBConnection"

	connStr := cfg.URL
	if connStr == "" {
		connStr = fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=%s",
			cfg.Host, cfg.Port, cfg.Username, cfg.DBName, cfg.Password, cfg.SSLMode)
	}
	db, err := sqlx.Connect("postgres", connStr)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}
	return &Storage{DB: db}, nil
}

func (s *Storage) Stop() error {
	return s.DB.Close()
}