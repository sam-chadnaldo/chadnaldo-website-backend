package main

import (
	"log/slog"
	"os"

	"github.com/sam-chadnaldo/chadnaldo-website-backend/internal/app"
	"github.com/sam-chadnaldo/chadnaldo-website-backend/internal/config"
	slogpretty "github.com/sam-chadnaldo/chadnaldo-website-backend/internal/lib/pretter"
)

func main() {

	cfg := config.MustLoad()

	log := setupLogger(cfg.Env)

	app := app.New(log, &cfg.Http)

	if err := app.Run(); err != nil{
		log.Error(
			"http server error",
			slog.Any("error", err),
		)
		os.Exit(1)
	}

}

func setupLogger(env string) *slog.Logger {
	var log *slog.Logger

	switch env {
	case config.EnvLocal:
		log = setupPrettySlog()
	case config.EnvDev:
		log = slog.New(
			slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}),
		)
	case config.EnvProd:
		log = slog.New(
			slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelInfo}),
		)
	}
	return log
}

func setupPrettySlog() *slog.Logger {
	opts := slogpretty.PrettyHandlerOptions{
		SlogOpts: &slog.HandlerOptions{
			Level: slog.LevelDebug,
		},
	}

	handler := opts.NewPrettyHandler(os.Stdout)

	return slog.New(handler)
}