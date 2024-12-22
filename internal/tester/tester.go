package tester

import (
	"net/http"
	"sync"
	"time"
	
	"github.com/HTM1000/stress-tester/internal/config"
)

type Result struct {
	StatusCode int           `json:"status_code"`
	Duration   time.Duration `json:"duration"`
}

type Report struct {
	TotalTime       time.Duration `json:"total_time"`
	TotalRequests   int           `json:"total_requests"`
	Successful200   int           `json:"successful_200"`
	StatusCounts    map[int]int   `json:"status_counts"`
	RequestsDetails []Result      `json:"requests_details,omitempty"`
}

func RunLoadTest(cfg *config.Config) (*Report, error) {
	results := make(chan Result, cfg.Requests)
	var wg sync.WaitGroup

	for i := 0; i < cfg.Concurrency; i++ {
		wg.Add(1)
		go func(workerID int) {
			defer wg.Done()
			for j := workerID; j < cfg.Requests; j += cfg.Concurrency {
				start := time.Now()
				resp, err := http.Get(cfg.URL)
				duration := time.Since(start)

				if err != nil {
					results <- Result{StatusCode: 0, Duration: duration}
					continue
				}

				results <- Result{StatusCode: resp.StatusCode, Duration: duration}
				resp.Body.Close()
			}
		}(i)
	}

	go func() {
		wg.Wait()
		close(results)
	}()

	return processResults(results, cfg)
}

func processResults(results chan Result, cfg *config.Config) (*Report, error) {
	statusCounts := make(map[int]int)
	var totalDuration time.Duration
	var detailedResults []Result

	for res := range results {
		statusCounts[res.StatusCode]++
		totalDuration += res.Duration
		if cfg.Detail {
			detailedResults = append(detailedResults, res)
		}
	}

	totalTime := totalDuration
	return &Report{
		TotalTime:       totalTime,
		TotalRequests:   cfg.Requests,
		Successful200:   statusCounts[200],
		StatusCounts:    statusCounts,
		RequestsDetails: detailedResults,
	}, nil
}
