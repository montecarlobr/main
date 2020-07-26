package postgres

import (
	"context"

	"github.com/FelipeUmpierre/stock-prices-watcher/internal/app/stock"
	"github.com/jmoiron/sqlx"
)

type (
	StockPriceWriter struct {
		db *sqlx.DB
	}

	StockPricesReader struct {
		db *sqlx.DB
	}
)

func NewStockPricesWriter(db *sqlx.DB) *StockPriceWriter {
	return &StockPriceWriter{db: db}
}

func (s *StockPriceWriter) BeginTransaction(ctx context.Context) (*sqlx.Tx, error) {
	return s.db.BeginTxx(ctx, nil)
}

func (s *StockPriceWriter) Commit(tx *sqlx.Tx) error   { return tx.Commit() }
func (s *StockPriceWriter) Rollback(tx *sqlx.Tx) error { return tx.Rollback() }

func (s *StockPriceWriter) Save(ctx context.Context, tx *sqlx.Tx, stck *stock.StockPrices) error {
	_, err := tx.NamedExecContext(ctx, `INSERT INTO 
		stock_prices 
	(stocks_id, open, high, low, close, volume, time_series) VALUES
	(:stocks_id, :open, :high, :low, :close, :volume, :time_series) 
	ON CONFLICT (stocks_id, time_series) DO NOTHING;
	`, stck)

	return err
}

func NewStockPricesReader(db *sqlx.DB) *StockPricesReader {
	return &StockPricesReader{db: db}
}

func (s *StockPricesReader) FindByStockID(id string) error {
	return nil
}
