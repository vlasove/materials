// Package task32 Написать собственную функцию Sleep.
package task32

import (
	"fmt"
	"runtime"
	"time"
)

// Sleep ...
func Sleep(t time.Duration) {
	done := make(chan struct{})
	timer := time.NewTimer(t)
	go func(done chan struct{}) {
		for {
			select {
			case <-timer.C:
				done <- struct{}{}
				runtime.Goexit()
			default:
				continue
			}
		}

	}(done)
	<-done
}

// Start ...
func Start() {
	fmt.Println("Message")
	Sleep(5 * time.Second)
	fmt.Println("Message after sleep")
}
