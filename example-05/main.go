package main

import (
	"fmt"
	"sync"
)

// Job is a custom type to represent a Job with an attribute called Id.
type Job struct {
	// Id is the identifier of the job.
	Id int
}

// worker is function that is spawned as a goroutine. It belongs to a pool of workers
// that are processing the Job Queue and are active until the jobQueue channel is clsoed.
func worker(wg *sync.WaitGroup, workerId int, jobQueue chan Job) {
	defer wg.Done()

	// when there is only one channel to watch upon, range could be used instead of select.
	// the loop remains active until the channel (jobQueue in this case) is closed.
	for job := range jobQueue {
		fmt.Printf("Worker %d is executing Job %d\n", workerId, job.Id)
	}
}

// dispatcher can be thought of as a utility function that spawns a number of goroutines as a pool
// and dispatches the jobs to be processed by these fixed number of goroutines from the pool.
func dispatcher(jobs []Job, workerCount int) {
	jobQueue := make(chan Job)
	wg := &sync.WaitGroup{}

	wg.Add(workerCount)
	for i := 0; i < workerCount; i++ {
		go worker(wg, i, jobQueue)
	}

	for _, job := range jobs {
		jobQueue <- job
	}

	// close the jobQueue channel to notify the worker to terminate their job processing loops and
	// to inturn terminate the goroutines.
	close(jobQueue)
	wg.Wait()
}

func main() {
	workerCount := 3
	jobs := make([]Job, 10)

	for index := range jobs {
		jobs[index].Id = index + 1
	}

	dispatcher(jobs, workerCount)
}
