// Package task25 Написать свою структуру счетчик,
// которая будет инкрементировать и выводить значения в конкурентной среде.
package task25

import (
	"fmt"
	"sync"
)

// Counter ...
type Counter struct {
	c  int
	mu *sync.RWMutex
}

// NewCounter ...
func NewCounter() *Counter {
	return &Counter{0, new(sync.RWMutex)}
}

// String ...
func (c *Counter) String() string {
	c.mu.RLock()
	defer c.mu.RUnlock()
	return fmt.Sprintf("%d", c.c)
}

// Add ...
func (c *Counter) Add() {
	c.mu.Lock()
	c.c++
	c.mu.Unlock()
}

// Start ...
func Start() {
	var wg sync.WaitGroup
	counter := NewCounter()
	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func(id int, wg *sync.WaitGroup) {
			counter.Add()
			if id%10 == 0 {
				fmt.Println("In goroutine val:", counter)
			}
			wg.Done()
		}(i, &wg)
	}
	wg.Wait()
	fmt.Println("Current val:", counter)
}
