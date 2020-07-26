package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/FelipeUmpierre/stock-prices-watcher/internal/app/gateway"
	"github.com/FelipeUmpierre/stock-prices-watcher/internal/app/stock"
	"github.com/FelipeUmpierre/stock-prices-watcher/internal/app/storage/postgres"
	"github.com/jmoiron/sqlx"
	"github.com/kelseyhightower/envconfig"
	_ "github.com/lib/pq"
)

type config struct {
	BaseURL string `split_words:"true"`
	APIKey  string `split_words:"true"`
	DB      string
}

func main() {
	log.Println("starting")
	defer log.Println("done")

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	if err := run(ctx); err != nil {
		panic(err)
	}
}

func run(ctx context.Context) error {
	// =============================================
	// Configuration
	// =============================================
	var cfg config
	if err := envconfig.Process("", &cfg); err != nil {
		return err
	}

	// =============================================
	// AlphaVantage
	// =============================================
	alpha := gateway.NewAlphaVantageClient(&http.Client{
		Timeout: 30 * time.Second,
	}, gateway.AlphaVantageConfig{
		BaseURL: cfg.BaseURL,
		APIKey:  cfg.APIKey,
	})

	db, err := sqlx.ConnectContext(ctx, "postgres", cfg.DB)
	if err != nil {
		return fmt.Errorf("failed to connect to db: %w", err)
	}

	stockPricesRepository := postgres.NewStockPricesWriter(db)

	// =============================================
	// Service
	// =============================================
	service := stock.NewStockService(alpha, stockPricesRepository)
	if err := service.CollectIntraday(ctx, "MSFT"); err != nil {
		return err
	}

	return nil
}
