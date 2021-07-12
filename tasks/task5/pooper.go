// Package task5 Написать программу,
// которая будет последовательно писать значения в канал,
// а с другой стороны канала — читать.
// По истечению N секунд программа должна завершиться.
package task5

import (
	"context"
	"fmt"
	"runtime"
	"time"
)

var (
	timeDelta = time.Second * 3
)

// Sender is function for sending data to channel
// and terminates if ctx.Done() will recieved
func Sender(ctx context.Context, num chan int) {
	v := 0
	for {
		select {
		case <-ctx.Done():
			runtime.Goexit()
		default:
			fmt.Printf("Sending %d to channel\n", v)
			num <- v
			v++
			time.Sleep(100 * time.Millisecond)
		}
	}
}

// Start ...
func Start() {
	numCh := make(chan int)
	ctx, cancelFunc := context.WithCancel(context.Background())
	go Sender(ctx, numCh)
	timer := time.NewTimer(timeDelta)

	for {
		select {
		case val := <-numCh:
			fmt.Printf("Recieve %d from channel\n", val)
		case <-timer.C:
			fmt.Println("Oh, time finished. Exited...")
			cancelFunc()
			return
		}
	}

}
