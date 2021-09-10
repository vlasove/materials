package main

import "fmt"

func main() {
	jobs := make(chan int, 5)
	done := make(chan bool)

	go func() {
		for {
			j, more := <-jobs
			if more {
				fmt.Println("recieved job:", j)
			} else {
				fmt.Println("recieved all jobs")
				done <- true
				return
			}
		}
	}()
	go func() {
		for j := 0; j < 4; j++ {
			jobs <- j
			fmt.Println("sent job:", j)
		}
		close(jobs)
		fmt.Println("sent all jobs")
	}()

	<-done
}
