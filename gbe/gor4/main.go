package main

import "fmt"

func ping(pings chan<- string, msg string) {
	pings <- msg
}

func pong(pings <-chan string, pongs chan<- string) {
	pongs <- <-pings
	pongs <- "hek"
}

func main() {
	pings := make(chan string, 1)
	pongs := make(chan string, 2)
	go ping(pings, "passed message")
	go pong(pings, pongs)
	fmt.Println(<-pongs, <-pongs)
}
