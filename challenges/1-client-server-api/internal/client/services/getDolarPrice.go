package services

import (
	"context"
	"encoding/json"
	"net/http"
	"time"
)

type DollarPriceResult struct {
	Bid string `json:"bid"`
}

func GetDolarPrice(ctx context.Context) (*DollarPriceResult, error) {
	timeoutCtx, cancel := context.WithTimeout(ctx, 300*time.Millisecond)
	defer cancel()

	req, err := http.NewRequestWithContext(timeoutCtx, http.MethodGet, "http://localhost:8080/cotacao", nil)
	if err != nil {
		return nil, err
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var response DollarPriceResult
	err = json.NewDecoder(resp.Body).Decode(&response)

	return &response, err
}
