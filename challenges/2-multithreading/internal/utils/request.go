package utils

import (
	"context"
	"io"
	"net/http"
	"time"
)

func RequestWithContext(ctx context.Context, url string) ([]byte, error) {
	ctx, cancel := context.WithTimeout(ctx, 1*time.Second)
	defer cancel()

	request, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}

	result, err := http.DefaultClient.Do(request)
	if err != nil {
		return nil, err
	}
	defer result.Body.Close()

	response, err := io.ReadAll(result.Body)
	if err != nil {
		return nil, err
	}

	return response, nil
}
