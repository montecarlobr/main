package postgres

import "database/sql"

type (
	StocksWriter struct {
		db *sql.DB
	}

	StocksReader struct {
		db *sql.DB
	}
)

func NewStocksWriter(db *sql.DB) *StocksWriter {
	return &StocksWriter{db: db}
}

func (s *StocksWriter) Save() error {
	return nil
}

func NewStocksReader(db *sql.DB) *StocksReader {
	return &StocksReader{db: db}
}

func (s *StocksReader) FindByID(id string) error {
	return nil
}
