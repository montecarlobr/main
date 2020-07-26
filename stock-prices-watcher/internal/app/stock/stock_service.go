package stock

import (
	"context"
	"fmt"
	"log"

	"github.com/FelipeUmpierre/stock-prices-watcher/internal/app/gateway"
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
)

type (
	AlphaVantage interface {
		GetDataFromStock(context.Context, string)
		Serie() <-chan *gateway.Serie
		Error() <-chan error
	}

	StockPriceRepository interface {
		BeginTransaction(context.Context) (*sqlx.Tx, error)
		Commit(*sqlx.Tx) error
		Rollback(*sqlx.Tx) error
		Save(context.Context, *sqlx.Tx, *StockPrices) error
	}

	Service struct {
		client AlphaVantage
		writer StockPriceRepository
	}
)

func NewStockService(client AlphaVantage, writer StockPriceRepository) *Service {
	return &Service{client: client, writer: writer}
}

func (s *Service) CollectIntraday(ctx context.Context, symbol string) error {
	go s.client.GetDataFromStock(ctx, symbol)

	tx, err := s.writer.BeginTransaction(ctx)
	if err != nil {
		return fmt.Errorf("failed to begin transaction: %w", err)
	}
	defer s.writer.Rollback(tx)

	go func() {
		for err := range s.client.Error() {
			log.Printf("failed: %v", err)
		}
	}()

	stockID := uuid.MustParse("e74f988d-cff0-44fb-8b82-f508e9cf846f")

	for serie := range s.client.Serie() {
		stock := NewStockPrices(uuid.New())
		stock.WithStocksID(stockID)
		stock.WithPrices(
			serie.Open,
			serie.High,
			serie.Low,
			serie.Close,
			serie.Volume,
		)
		stock.WithTimeSeries(serie.Time)

		if err := s.writer.Save(ctx, tx, stock); err != nil {
			return fmt.Errorf("failed to save: %w", err)
		}
	}

	if err := s.writer.Commit(tx); err != nil {
		return fmt.Errorf("failed to commit: %w", err)
	}

	return nil
}
