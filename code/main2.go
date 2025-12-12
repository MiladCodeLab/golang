//package main
//
//import (
//	"fmt"
//	"log"
//	"sync"
//	"time"
//)
//
//// ListQuery represents a job/query
//type ListQuery struct {
//	ID int
//}
//
//// ResponseType represents a job result
//type ResponseType struct {
//	QueryID int
//	Result  string
//}
//
//// httpCall simulates an HTTP call
//func httpCall(q ListQuery) (ResponseType, error) {
//	// Simulate some latency
//	time.Sleep(time.Duration(100+q.ID*10) * time.Millisecond)
//	fmt.Printf("Processed query ID: %d\n", q.ID)
//	return ResponseType{QueryID: q.ID, Result: "ok"}, nil
//}
//
//// ProcessJobs runs jobs concurrently using plain goroutines
//func ProcessJobs(jobsQuery []ListQuery) []ResponseType {
//	var wg sync.WaitGroup
//	var mu sync.Mutex
//
//	responses := make([]ResponseType, 0, len(jobsQuery))
//
//	for _, q := range jobsQuery {
//		wg.Add(1)
//		go func(job ListQuery) { // TODO: Worker Pool  4 kbyte
//			defer wg.Done()
//
//			// Execute the HTTP call
//			res, err := httpCall(job)
//			if err != nil {
//				log.Printf("httpCall error: %v", err)
//				return
//			}
//
//			// Safely append the result
//			mu.Lock()
//			responses = append(responses, res) // TODO: using channel instead
//			mu.Unlock()
//		}(q)
//	}
//
//	wg.Wait() // wating point
//	return responses
//}
//
//func main() {
//	var count int = 1000
//	// Simulate 1000 jobs
//	jobsQuery := make([]ListQuery, count)
//	for i := 0; i < count; i++ {
//		jobsQuery[i] = ListQuery{ID: i + 1}
//	}
//
//	results := ProcessJobs(jobsQuery)
//
//	fmt.Println("All results:")
//	for _, r := range results {
//		fmt.Printf("QueryID=%d, Result=%s\n", r.QueryID, r.Result)
//	}
//}
