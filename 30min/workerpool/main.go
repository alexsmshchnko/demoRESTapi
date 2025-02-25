package main

import (
	"sync"
	"time"
)

func worker(id int, jobs <-chan int, results chan<- int, wg *sync.WaitGroup) {
	for j := range jobs {
		println("worker", id, ": started ", j)
		time.Sleep(1 * time.Second)
		println("worker", id, ": finished ", j)
		results <- j
	}
	wg.Done()
}

const (
	MAX_WORKERS = 3
	NUM_JOBS    = 7
)

func main() {
	jobs := make(chan (int))
	wgJobs := sync.WaitGroup{}

	results := make(chan (int), NUM_JOBS)
	wgWorkers := sync.WaitGroup{}
	go func() {
		for r := range results {
			println(r)
		}
	}()

	wgWorkers.Add(MAX_WORKERS)
	for i := range MAX_WORKERS {
		go worker(i, jobs, results, &wgWorkers)
	}

	wgJobs.Add(NUM_JOBS)
	go func(wg *sync.WaitGroup) {
		for i := range NUM_JOBS {
			jobs <- i
			wg.Done()
		}
	}(&wgJobs)

	wgJobs.Wait()
	close(jobs)

	wgWorkers.Wait()
	close(results)

	time.Sleep(1 * time.Second)
}
