package loadtester

import (
	"net/http"
	"sync"
	"time"
)


func worker(id int, jobs <-chan Job, results chan<- Result, wg *sync.WaitGroup) {
	defer wg.Done()

	for job := range jobs {
		start := time.Now()
		
		// Make the request
		resp, err := http.Get(job.URL)
		
		duration := time.Since(start)

		result := Result{
			JobID:    job.ID,
			Duration: duration,
			Error:    err,
		}

		if err == nil {
			result.StatusCode = resp.StatusCode
			resp.Body.Close() 
		}

		// Send the result
		results <- result
	}
}