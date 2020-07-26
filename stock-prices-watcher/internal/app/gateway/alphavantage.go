package gateway

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"time"
)

const FunctionTimeSeriesIntraday = "TIME_SERIES_INTRADAY"

type (
	httpClient interface {
		Do(*http.Request) (*http.Response, error)
	}

	AlphaVantageConfig struct {
		BaseURL string
		APIKey  string
	}

	Alpha struct {
		client httpClient
		cfg    AlphaVantageConfig
		serie  chan *Serie
		err    chan error
	}
)

func NewAlphaVantageClient(client httpClient, cfg AlphaVantageConfig) *Alpha {
	return &Alpha{
		client: client,
		cfg:    cfg,
		serie:  make(chan *Serie),
		err:    make(chan error),
	}
}

func (a *Alpha) GetDataFromStock(ctx context.Context, symbol string) {
	u, _ := url.Parse(a.cfg.BaseURL)

	v := url.Values{}
	v.Add("function", FunctionTimeSeriesIntraday)
	v.Add("symbol", symbol)
	v.Add("interval", "1min")
	v.Add("apikey", a.cfg.APIKey)

	u.RawQuery = v.Encode()

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, u.String(), nil)
	if err != nil {
		a.err <- fmt.Errorf("failed to create the request: %w", err)
	}

	resp, err := a.client.Do(req)
	if err != nil {
		a.err <- fmt.Errorf("failed to execute the request: %w", err)
	}
	defer resp.Body.Close()

	dec := json.NewDecoder(resp.Body)

	dec.Token()

	for dec.More() {
		dec.Token()
	}

	dec.Token()
	dec.Token()
	dec.Token()

	for dec.More() {
		t, err := dec.Token()
		if err != nil {
			a.err <- fmt.Errorf("failed to get the token: %w", err)
		}

		tt, err := time.Parse("2006-01-02 15:04:05", t.(string))
		if err != nil {
			a.err <- fmt.Errorf("failed to parse time: %w", err)
		}

		i := Serie{Time: tt}
		if err := dec.Decode(&i); err != nil {
			a.err <- fmt.Errorf("failed to decode json: %w", err)
		}

		a.serie <- &i
	}

	close(a.serie)
}

func (a *Alpha) Serie() <-chan *Serie { return a.serie }
func (a *Alpha) Error() <-chan error  { return a.err }
