package main

import (
	"fmt"
	"time"
)

const (
	numOfJobs = 10
)

func worker(id int, jobs <-chan int, results chan<- int) {
	for j := range jobs {
		fmt.Println("worker", id, "started job", j)
		time.Sleep(time.Second)
		fmt.Println("worked", id, "finished job", j)
		results <- j * 4
	}
}

func main() {
	jobs := make(chan int, numOfJobs)
	results := make(chan int, numOfJobs)

	for w := 1; w <= 3; w++ {
		go worker(w, jobs, results)
	}

	for j := 0; j < numOfJobs; j++ {
		jobs <- j
	}
	close(jobs)

	for a := 1; a <= numOfJobs; a++ {
		fmt.Println(<-results)
	}

}
