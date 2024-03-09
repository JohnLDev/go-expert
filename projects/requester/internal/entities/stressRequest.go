package entities

import "fmt"

type StressRequest struct {
	Url         string
	Requests    int
	Concurrency int
}

type StressResult struct {
	Success            int
	Failed             int
	Requests           int
	Time               int64
	StatusDistribution map[int]int
}

func (s *StressResult) Present() {
	fmt.Println("x===================== Stress test result =====================x")
	fmt.Printf("Amount of requests with http status 200: %d\n", s.Success)
	fmt.Printf("Amount of failed requests: %d\n", s.Failed)
	fmt.Printf("Total amount of requests: %d\n", s.Requests)

	fmt.Printf("Total time: %.2fs\n", float64(s.Time)/1000)
	fmt.Println("Status distribution:")
	for k, v := range s.StatusDistribution {
		fmt.Printf("	Status %d: %d\n", k, v)
	}
	fmt.Println("x===================== Stress test result =====================x")

}
