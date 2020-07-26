package stock

import (
	"time"

	"github.com/google/uuid"
)

type (
	StockPrices struct {
		ID         uuid.UUID `json:"-" db:"id"`
		StocksID   uuid.UUID `json:"stocks_id" db:"stocks_id"`
		Open       string    `json:"open" db:"open"`
		High       string    `json:"high" db:"high"`
		Low        string    `json:"low" db:"low"`
		Close      string    `json:"close" db:"close"`
		Volume     string    `json:"volume" db:"volume"`
		TimeSeries time.Time `json:"time_series" db:"time_series"`
	}
)

// NewStockPrices creates a new stock price
func NewStockPrices(id uuid.UUID) *StockPrices { return &StockPrices{ID: id} }

// WithStocksID adds the stocks id
func (s *StockPrices) WithStocksID(id uuid.UUID) { s.StocksID = id }

// WithPrices adds the prices
func (s *StockPrices) WithPrices(open, high, low, close, vlm string) {
	s.Open = open
	s.High = high
	s.Low = low
	s.Close = close
	s.Volume = vlm
}

// WithTimeSeries adds the time series
func (s *StockPrices) WithTimeSeries(t time.Time) { s.TimeSeries = t }
