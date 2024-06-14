package hw05parallelexecution

import (
	"errors"
)

var ErrErrorsLimitExceeded = errors.New("errors limit exceeded")
var ErrNoWorkers = errors.New("no workers available")

type Task func() error

// Run starts tasks in n goroutines and stops its work when receiving m errors from tasks.
func Run(tasks []Task, n, m int) error {
	activeJobs := 0
	i := 0
	errs := 0
	taskLen := len(tasks)
	tasksChan := make(chan int, n)
	if n <= 0 {
		return ErrNoWorkers
	}
	for {
		if i >= taskLen {
			break
		}
		if activeJobs < n {
			if m > 0 && errs >= m {
				return ErrErrorsLimitExceeded
			}
			task := tasks[i]
			go func() {
				defer func() {
					tasksChan <- 0
				}()
				err := doOneJob(task)
				if err != nil {
					errs++
				}
			}()
			activeJobs++
			i++
		}
		<-tasksChan
		activeJobs--
	}
	return nil
}

func doOneJob(fn Task) error {
	return fn()
}
