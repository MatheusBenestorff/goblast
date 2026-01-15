package loadtester

import (
	"fmt"
	"sync"
	"time"
)

func RunLoadTest(url string, totalRequests int, concurrency int) {
	fmt.Printf("Iniciando ataque a %s\n", url)

	jobs := make(chan Job, totalRequests)
	results := make(chan Result, totalRequests)

	// WaitGroup for Workers
	var wg sync.WaitGroup

	// Start the Workers
	for w := 1; w <= concurrency; w++ {
		wg.Add(1)
		go worker(w, jobs, results, &wg)
	}

	// Start Metrics Collector
	done := make(chan bool)
	go func() {
		processResults(results)
		done <- true
	}()

	// Send the Jobs (Producer)
	start := time.Now()
	for i := 1; i <= totalRequests; i++ {
		jobs <- Job{URL: url, ID: i}
	}
	close(jobs) 

	// Wait all workers
	wg.Wait()

	// Close results channel
	close(results)

	<-done

	elapsed := time.Since(start)
	fmt.Printf("\nCarga finalizada em %s\n", elapsed)
}

func processResults(results <-chan Result) {
	var (
		totalDuration time.Duration
		successCount  int
		errorCount    int
	)

	// Read from the channel until it is closed
	for res := range results {
		if res.Error != nil {
			errorCount++
			// fmt.Printf("Erro no Job %d: %v\n", res.JobID, res.Error)
		} else {
			successCount++
			totalDuration += res.Duration
		}
	}

	// Calculating simple average
	avg := time.Duration(0)
	if successCount > 0 {
		avg = totalDuration / time.Duration(successCount)
	}

	fmt.Println("\nRelatório Final:")
	fmt.Printf("Sucessos: %d\n", successCount)
	fmt.Printf("Erros:    %d\n", errorCount)
	fmt.Printf("Média:    %s\n", avg)
}