package app

import (
	"context"
	"log/slog"

	"github.com/sam-chadnaldo/chadnaldo-website-backend/internal/app/server"
	"github.com/sam-chadnaldo/chadnaldo-website-backend/internal/config"
)

type App struct {
	log    *slog.Logger
	server *server.Server
	port string
}

func New(
	log *slog.Logger, 
	cfg *config.HttpConfig,

	) *App {
	router := server.NewHandler(log)
	routes := router.InitRouts()
	srv := server.NewServer(cfg.Port, routes)

	return &App{
		log:    log,
		server: srv,
		port: cfg.Port,
	}
}

func (a *App) Run() error {
	a.log.Info("Starting the application", slog.String("port", a.port))
	return a.server.Run()
}

func (a *App) GracefulShutdown(ctx context.Context) error {
	a.log.Info("Initiating graceful shutdown")
	return a.server.Shutdown(ctx)
}