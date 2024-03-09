package usecases

import (
	"fmt"
	"net/http"
	"slices"
	"sync"
	"time"

	"github.com/johnldev/requester/internal/entities"
)

type StressTestUseCase struct {
}

func (uc *StressTestUseCase) Execute(input entities.StressRequest) (entities.StressResult, error) {
	var initialTime = time.Now()

	fmt.Printf("Input: %+v\n", input)

	result := entities.StressResult{
		StatusDistribution: make(map[int]int),
		Success:            0,
		Failed:             0,
		Requests:           0,
		Time:               0,
	}

	wg := sync.WaitGroup{}
	wg.Add(input.Requests)

	channel := make(chan int, input.Concurrency)
	defer close(channel)

	for i := 0; i < input.Concurrency; i++ {
		go func() {
			for range channel {
				response, err := http.DefaultClient.Get(input.Url)

				if err != nil {
					fmt.Printf("Error: %s\n", err.Error())
					result.Failed++
				} else {
					badStatusCodes := []int{
						http.StatusBadRequest,
						http.StatusBadGateway,
						http.StatusNotFound,
						http.StatusInternalServerError,
						http.StatusUnauthorized,
						http.StatusForbidden,
					}
					if !slices.Contains(badStatusCodes, response.StatusCode) {
						result.Success++
					} else {
						result.Failed++
					}
					result.StatusDistribution[response.StatusCode]++
				}
				result.Requests++
				wg.Done()
			}
		}()
	}

	for i := 0; i < input.Requests; i++ {
		channel <- i
	}

	wg.Wait()

	result.Time = time.Since(initialTime).Milliseconds()

	return result, nil
}

func InitStressTestUseCase() *StressTestUseCase {
	return &StressTestUseCase{}
}
