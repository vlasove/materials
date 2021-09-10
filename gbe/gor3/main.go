package main

import (
	"fmt"
	"time"
)

func worker(done chan bool) {
	fmt.Print("working....")
	time.Sleep(time.Second)
	fmt.Println("done")
	done <- true
}

func main() {
	done := make(chan bool)
	go worker(done)
	for {
		select {
		case <-done:
			fmt.Println("alles gut!")
			return
		default:
			fmt.Println("waiting!!")
			time.Sleep(time.Millisecond * 100)
		}
	}
}
