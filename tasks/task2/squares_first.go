// Package task2 Написать программу,
// которая конкурентно рассчитает
// значение квадратов значений взятых из массива
// (2,4,6,8,10) и выведет их квадраты в stdout.
package task2

import (
	"fmt"
	"sync"
)

var (
	arr = [5]int{2, 4, 6, 8, 10}
)

// Start ...
func Start() {
	var wg sync.WaitGroup
	for id, val := range arr {
		wg.Add(1)
		// send copy of id and val + ptr to wg
		go func(id int, val int, wg *sync.WaitGroup) {
			arr[id] = val * val
			wg.Done()
		}(id, val, &wg)

	}
	wg.Wait()
	fmt.Println(arr)
}
