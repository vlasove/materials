// Package task3 Дана
// последовательность чисел  (2,4,6,8,10)
// найти их сумму квадратов(22+32+42….) с использованием конкурентных вычислений.
package task3

import (
	"sync"
)

var (
	arr = []int{2, 4, 6, 8, 10}
)

// StartWG is function with mutexes and wait groups
func StartWG() int {
	var wg sync.WaitGroup
	var mu sync.RWMutex
	sum := 0
	for _, v := range arr {
		wg.Add(1)
		// sends copy of val in function
		go func(val int, mu *sync.RWMutex, wg *sync.WaitGroup) {
			// block for writing
			mu.Lock()
			sum += val * val
			mu.Unlock()

			wg.Done()
		}(v, &mu, &wg)

	}
	wg.Wait()
	return sum
}
