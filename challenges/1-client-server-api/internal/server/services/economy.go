package services

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

type PriceResult struct {
	Code       string `json:"code"`
	Codein     string `json:"codein"`
	Name       string `json:"name"`
	High       string `json:"high"`
	Low        string `json:"low"`
	VarBid     string `json:"varBid"`
	PctChange  string `json:"pctChange"`
	Bid        string `json:"bid"`
	Ask        string `json:"ask"`
	Timestamp  string `json:"timestamp"`
	CreateDate string `json:"create_date"`
}

type DollarPriceResult struct {
	Usdbrl PriceResult `json:"USDBRL"`
}

func GetDollarPrice(ctx context.Context) (*DollarPriceResult, error) {
	timeoutCtx, cancel := context.WithTimeout(ctx, 200*time.Millisecond)
	defer cancel()

	request, err := http.NewRequestWithContext(timeoutCtx, http.MethodGet, "https://economia.awesomeapi.com.br/json/last/USD-BRL", nil)
	if err != nil {
		return nil, err
	}

	resp, err := http.DefaultClient.Do(request)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	var response DollarPriceResult
	err = json.NewDecoder(resp.Body).Decode(&response)
	if err != nil {
		return nil, err
	}

	fmt.Printf("Dollar price: %+v\n", response)

	return &response, nil
}
