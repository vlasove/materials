// Package task9 Написать конвейер чисел.
// Даны 2 канала - в первый пишутся числа из массива,
// во второй пишется результат операции 2*x,
// после чего данные выводятся в stdout.
package task9

import (
	"fmt"
	"io"
	"os"
)

var (
	arr = [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
)

// Sender ...
func Sender(arr [10]int, sendChan chan<- int) {
	for _, v := range arr {
		sendChan <- v
	}
	close(sendChan)
}

// Reciever ...
func Reciever(sendChan <-chan int, recvChan chan<- int) {
	for v := range sendChan {
		recvChan <- v * 2
	}
	close(recvChan)
}

// Printer ...
func Printer(w io.Writer, recvChan chan int, done chan<- struct{}) {
	for v := range recvChan {
		fmt.Fprintf(w, "%d ", v)
	}
	fmt.Fprintf(w, "\n")
	done <- struct{}{}
}

// Start ...
func Start() {
	sendChan := make(chan int)
	outChan := make(chan int)

	done := make(chan struct{})

	go Sender(arr, sendChan)
	go Reciever(sendChan, outChan)
	go Printer(os.Stdout, outChan, done)
	<-done

}
