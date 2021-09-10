package main

import "fmt"

func main() {
	messages := make(chan string)

	go func() {
		messages <- "ping"
	}()

	msg, ok := <-messages
	if ok {
		fmt.Println(msg)
	}
}
