package handlers

import (
	"log"

	"github.com/johnldev/requester/internal/entities"
	usecases "github.com/johnldev/requester/internal/useCases"
)

func StressTestHandler(url string, requests, concurrency int) {
	input := entities.StressRequest{
		Url:         url,
		Requests:    requests,
		Concurrency: concurrency,
	}

	useCase := usecases.InitStressTestUseCase()
	res, err := useCase.Execute(input)
	if err != nil {
		log.Fatal(err)
	}
	res.Present()
}
