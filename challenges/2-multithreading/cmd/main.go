package main

import (
	"context"

	"github.com/johnldev/go-expert/chanllenge-2/internal/services"
)

func main() {
	viaCepService := services.NewCepService(context.Background())
	viaCepService.GetCep("01153-000")
}
