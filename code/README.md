# Concurrency Interview Scenario

---

## **Scenario & Constraints for the Interview**

**Scenario:**
Your service receives multiple user requests. For each request, the service needs to perform several side jobs (e.g., HTTP calls to external APIs) to collect additional data before preparing the main response. Each job is independent, and we want to process them concurrently to reduce total runtime.

**Input:**

* A slice of jobs (`[]ListQuery`), each representing a unit of work.
* The number of jobs per request can vary (e.g., 10–20 per user request).
* Total requests can arrive simultaneously (e.g., 1,000 concurrent requests).

---

### **Constraints / Edge Cases for Analysis**

1. **Concurrency Safety:**

    * Multiple goroutines will try to append results to a shared slice (`responses`) simultaneously.
    * Risk: race conditions, corrupted slice, or panics.
    * Solution in naive version: `sync.Mutex` to protect the shared slice.

2. **High Concurrency / Resource Limits:**

    * Launching thousands of goroutines at once may:
        * Consume excessive memory (goroutine stacks).
        * Exhaust file descriptors or network connections.
        * Overload CPU scheduling.
        * Risk OS termination in extreme cases.
    * Constraint to analyze: how to **limit concurrency safely**.

3. **Job Failures:**

    * Some jobs may fail (network errors, invalid responses).
    * Constraint: failure in one job **should not stop other jobs**.
    * Collect as many results as possible from successful jobs.

4. **Variable Job Duration / Latency:**

    * Jobs can take different times (some slow, some fast).
    * Constraint: **long-running jobs should not block overall processing**.

5. **Order of Results:**

    * `responses` slice may receive results **out of order**, depending on job completion times.
    * Constraint: candidate should consider if ordering matters for the system.

6. **Resource Cleanup:**

    * Goroutines that finish must release any resources (e.g., locks, memory).
    * Constraint: ensure **no goroutine leaks** or deadlocks.

7. **System-wide Resource Guarantees (Optional, Advanced):**

    * If worker pool or semaphore were **global across all requests**, max concurrency can be controlled at system level.
    * Constraint: candidate should consider **per-request vs global concurrency limits**.

---

### **Interview Objective**

Ask the candidate to:

1. **Identify risks and constraints** in the naive version.
2. **Propose improvements** for:

    * Resource usage (CPU, memory, network).
    * Safe concurrency control (worker pool, semaphore, context).
    * Handling failures and timeouts.
3. **Explain trade-offs** between simplicity, safety, and performance.

---

## **Sample Go Code (Naive Version for Interview)**

```go
package main

import (
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
func httpCall(q ListQuery) (ResponseType, error) {
	// Simulate some latency
	time.Sleep(time.Duration(100+q.ID*10) * time.Millisecond)
	fmt.Printf("Processed query ID: %d\n", q.ID)
	return ResponseType{QueryID: q.ID, Result: "ok"}, nil
}

// ProcessJobs runs jobs concurrently using plain goroutines
func ProcessJobs(jobsQuery []ListQuery) []ResponseType {
	var wg sync.WaitGroup
	var mu sync.Mutex

	responses := make([]ResponseType, 0, len(jobsQuery))

	for _, q := range jobsQuery {
		wg.Add(1)
		go func(job ListQuery) {
			defer wg.Done()

			// Execute the HTTP call
			res, err := httpCall(job)
			if err != nil {
				log.Printf("httpCall error: %v", err)
				return
			}

			// Safely append the result
			mu.Lock()
			responses = append(responses, res)
			mu.Unlock()
		}(q)
	}

	wg.Wait()
	return responses
}

func main() {
	var count int = 1000
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
```

