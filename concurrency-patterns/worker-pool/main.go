package main

import (
	"fmt"
	"sync"
	"time"
)

func worker(id int, jobs <-chan int, results chan<- string) {
	fmt.Println(id, " worked id is working")
	for job := range jobs {
		fmt.Printf("Worker %d processing job %d\n", id, job)
		results <- fmt.Sprintf("%d job is signed by by %d", job, id)
		time.Sleep(1 * time.Second)
	}
}

func main() {
	numJobs := 10
	numWorkers := 3

	jobs := make(chan int)
	results := make(chan string, numJobs)

	var wg sync.WaitGroup

	// Start worker goroutines
	for i := 1; i <= numWorkers; i++ {
		wg.Add(1)
		go func(workerID int) {
			defer wg.Done()
			worker(workerID, jobs, results)
		}(i)
	}

	// Enqueue jobs
	for i := 1; i <= numJobs; i++ {
		jobs <- i
	}
	close(jobs)

	// Wait for all workers to finish
	go func() {
		wg.Wait()
		close(results)
	}()

	finalResult := ""
	// Collect results
	for result := range results {
		finalResult = fmt.Sprintf(finalResult, result)
	}

	fmt.Println(finalResult)
}
