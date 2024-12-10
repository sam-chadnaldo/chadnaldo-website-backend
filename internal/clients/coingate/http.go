package coingate

import (
	"log/slog"
	"net/http"
	"time"

	"github.com/sam-chadnaldo/chadnaldo-website-backend/internal/config"
)

type Client struct {
	client 			http.Client
	log 			slog.Logger
	baseURL			string
	apiToken			string
	timeout 		time.Duration
	retriesCount 	int
}

func New(log slog.Logger, cfg config.CoinGateConfig, env string) *Client {
	client := http.Client{}

	var url string

	if env == config.EnvLocal || env == config.EnvDev{
		url = cfg.TestURL
	} else {
		url = cfg.BaseURL
	}

	return &Client{
		client: client,
		log: log,
		baseURL: url,
		apiToken: cfg.ApiToken,
		timeout: cfg.Timeout,
		retriesCount: cfg.RetriesCount,

	}
}