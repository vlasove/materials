// Package task33 Даны 2 канала - в первый пишутся рандомные числа после чего они проверяются
// на четность и отправляются во второй канал.
// Результаты работы из второго канала пишутся в stdout.
package task33

import (
	"fmt"
	"io"
	"math/rand"
	"os"
)

// Sender ...
func Sender(n int, sendChan chan<- int) {
	for i := 0; i < n; i++ {
		sendChan <- rand.Intn(100500)
	}
	close(sendChan)
}

// EvenReciever ...
func EvenReciever(recvChan <-chan int, sendChan chan<- *CompositeValue) {
	for v := range recvChan {
		sendChan <- NewCompositeValue(v)
	}
	close(sendChan)
}

// Printer ...
func Printer(w io.Writer, recvChan <-chan *CompositeValue, done chan struct{}) {
	for cv := range recvChan {
		fmt.Fprintf(w, "Value:%d Even?:%v\n", cv.val, cv.isEven)
	}
	done <- struct{}{}
}

// CompositeValue ...
type CompositeValue struct {
	val    int
	isEven bool
}

// NewCompositeValue ...
func NewCompositeValue(val int) *CompositeValue {
	return &CompositeValue{val, val%2 == 0}
}

// Start ...
func Start() {
	sendChan := make(chan int)
	recvChan := make(chan *CompositeValue)
	done := make(chan struct{})
	go Sender(10, sendChan)
	go EvenReciever(sendChan, recvChan)
	go Printer(os.Stdout, recvChan, done)
	<-done
}
