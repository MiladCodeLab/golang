package main

import (
	"context"
	"fmt"
	"log"
	"sync"
	"time"
)

// ListQuery represents a job/query
type ListQuery struct {
	ID int
}

// ResponseType represents a job result
type ResponseType struct {
	QueryID int
	Result  string
}

// httpCall simulates an HTTP call
func httpCall(ctx context.Context, q ListQuery) (ResponseType, error) {
	ctxTimeout, cancel := context.WithTimeout(ctx, 30*time.Second)
	defer cancel()
	// Simulate some latency
	time.Sleep(time.Duration(100+q.ID*10) * time.Millisecond)
	fmt.Printf("Processed query ID: %d\n", q.ID)
	return ResponseType{QueryID: q.ID, Result: "ok"}, nil
}

const (
	WorkersNum int = 10
)

// ProcessJobs runs jobs concurrently using plain goroutines
func ProcessJobs(ctx context.Context, jobsQuery []ListQuery) []ResponseType {
	var wg sync.WaitGroup
	var mu sync.Mutex
	input := make(chan ListQuery)

	responses := make([]ResponseType, 0, len(jobsQuery))
	// goroutine sending jobs to worker pool
	wg.Go(func() {
		for _, q := range jobsQuery {
			<-ctx.Done()
			input <- q // insert into worker pool
		}
		close(input)
	})

	// worker pool
	for i := 0; i < WorkersNum; i++ {
		wg.Go(func() {
			for {
				select {
				case <-ctx.Done():
					return
				case job := <-input:
					// Execute the HTTP call
					// heavy calculation
					//time.Sleep(time.Second * 3)
					var (
						res ResponseType
						err error
					)
					//maxHttpCh := make(,3)
					// max http connection == 3
					{
						res, err = httpCall(ctx, job)
						if err != nil {
							log.Printf("httpCall error: %v", err)
							return
						}
					}

					// Safely append the result
					mu.Lock()
					responses = append(responses, res) // TODO: using channel instead
					mu.Unlock()
				}
			}
		})
	}

	wg.Wait() // wating point
	return responses
}

func main() {
	var count int = 1_000
	// Simulate 1000 jobs
	jobsQuery := make([]ListQuery, count)
	for i := 0; i < count; i++ {
		jobsQuery[i] = ListQuery{ID: i + 1}
	}

	results := ProcessJobs(jobsQuery)

	fmt.Println("All results:")
	for _, r := range results {
		fmt.Printf("QueryID=%d, Result=%s\n", r.QueryID, r.Result)
	}
}
