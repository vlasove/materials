// Package task6 Какие существуют способы остановить
// выполнения горутины?
// Написать примеры использования.
package task6

import (
	"context"
	"fmt"
	"runtime"
	"time"
)

// StopWithChan ...
func StopWithChan(quit chan struct{}) {
	for {
		select {
		case <-quit:
			fmt.Println("StopWithChan done")
			runtime.Goexit()
		default:
			time.Sleep(500 * time.Millisecond)
			fmt.Println("StopWithChan working")
		}
	}
}

// StopWithTimer ...
func StopWithTimer(timer *time.Timer) {
	for {
		select {
		case <-timer.C:
			fmt.Println("StopWithTimer done")
			runtime.Goexit()
		default:
			time.Sleep(600 * time.Millisecond)
			fmt.Println("StopWithTime working")
		}
	}
}

// StopWithContext ...
func StopWithContext(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			fmt.Println("StopWithContext done")
			runtime.Goexit()
		default:
			time.Sleep(700 * time.Millisecond)
			fmt.Println("StopWithContext working")
		}
	}
}

// HowManyGoroutine ...
func HowManyGoroutine() {
	fmt.Printf("\n\nCurrently works %d goroutines\n\n", runtime.NumGoroutine())
}

// Start ...
func Start() {
	quit := make(chan struct{})
	ctx, cancel := context.WithCancel(context.Background())
	timer := time.NewTimer(11 * time.Second)
	go StopWithChan(quit)
	go StopWithContext(ctx)
	go StopWithTimer(timer)
	defer HowManyGoroutine()
	{
		HowManyGoroutine()
		time.Sleep(time.Second * 4)
		fmt.Println("Killing StopWithChan...")
		quit <- struct{}{}
		time.Sleep(time.Second * 1)
		HowManyGoroutine()

		time.Sleep(time.Second * 4)
		fmt.Println("Killing StopWithContext...")
		cancel()
		time.Sleep(time.Second * 1)
		HowManyGoroutine()

		time.Sleep(time.Second * 4)
		fmt.Println("Killing StopWithTimer... ")
		timer.Stop()
		time.Sleep(4 * time.Second)
		HowManyGoroutine()
	}

	fmt.Println("All goroutines killed")
}
