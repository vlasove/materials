package main

import (
	"fmt"
	"time"
)

func pinger(c chan<- string) {

	for i := 0; ; i++ {
		c <- "ping"
	}

}

func ponger(c chan<- string) {
	for i := 0; ; i++ {
		c <- "pong"
	}
}

func printer(c <-chan string) {
	for {
		msg := <-c
		fmt.Println(msg)
		time.Sleep(200 * time.Millisecond)
	}
}

func main() {
	c := make(chan string)

	go pinger(c)
	go ponger(c)
	go printer(c)

	time.Sleep(15 * time.Second)

	c1 := make(chan string)
	c2 := make(chan string)

	go func() {
		c1 <- "poop"
		time.Sleep(2 * time.Second)
	}()

	go func() {
		c2 <- "boop"
		time.Sleep(3 * time.Second)
	}()

	go func() {
	out:
		for {
			select {
			case msg := <-c1:
				fmt.Println(msg)
				break out
			case msg := <-c2:
				fmt.Println(msg)
				break out
			}
		}
		fmt.Println("Finish with select")
	}()
	time.Sleep(4 * time.Second)
}
