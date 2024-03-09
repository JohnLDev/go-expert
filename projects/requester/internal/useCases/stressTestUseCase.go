package usecases

import (
	"fmt"
	"net/http"
	"slices"
	"sync"

	"github.com/johnldev/requester/internal/entities"
	"github.com/johnldev/requester/internal/utils"
)

type StressTestUseCase struct {
}

func (uc *StressTestUseCase) Execute(input entities.StressRequest) (entities.StressResult, error) {
	fmt.Printf("Input: %+v\n", input)
	var timer = utils.StartTimer()

	result := entities.StressResult{
		StatusDistribution: make(map[int]int),
		Success:            0,
		Failed:             0,
		Requests:           0,
		Time:               0,
	}

	wg := sync.WaitGroup{}
	wg.Add(input.Requests)

	mutex := sync.Mutex{}

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

					mutex.Lock()
					result.StatusDistribution[response.StatusCode]++
					mutex.Unlock()

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
	result.Time = timer.StopTimer().Result
	return result, nil
}

func InitStressTestUseCase() *StressTestUseCase {
	return &StressTestUseCase{}
}
