package loadtester

import "time"

// Job to be done
type Job struct {
	URL string
	ID  int
}

// Request result
type Result struct {
	JobID      int
	StatusCode int
	Duration   time.Duration 
	Error      error
}