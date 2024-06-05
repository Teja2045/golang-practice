package main

import (
	"fmt"
	"time"
)

func worker(jobs <-chan int, done func()) {

	fmt.Println("worker started...")
	for job := range jobs {
		fmt.Println(job, " job is processing...")
	}
	done()
	fmt.Println("worked ended !")
}

func done() {
	fmt.Println("caller ending..............")
}

func main() {
	fmt.Println("testing....")

	jobs := make(chan int)
	go worker(jobs, done)

	for i := 0; i < 100; i++ {
		jobs <- i
	}

	time.Sleep(1 * time.Second)
	close(jobs)
	time.Sleep(time.Second)
}
