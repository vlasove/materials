// Package task21 Написать программу,
// которая в конкурентном виде читает элементы из массива в stdout.
package task21

import (
	"fmt"
	"sync"
)

var (
	arr = [7]int{1, 2, 3, 4, 5, 6, 7}
)

// Printer ...
// func Printer(w io.Writer, val int, wg *sync.WaitGroup) {
// 	fmt.Fprintf(w, "%d ", val)
// 	wg.Done()
// }

// // Start ...
// func Start() {
// 	var wg sync.WaitGroup

// 	for _, v := range arr {
// 		wg.Add(1)
// 		go Printer(os.Stdout, v, &wg)
// 	}
// 	wg.Wait()
// 	fmt.Println()
// }

// Start2 ...
func Start2() {
	wg := new(sync.WaitGroup)
	fmt.Println("Concurrent read from arr:")
	for i := 0; i < len(arr); i++ {
		wg.Add(1)
		go func(id int, wg *sync.WaitGroup) {
			fmt.Printf("%d ", arr[id])
			wg.Done()
		}(i, wg)

	}
	wg.Wait()
	fmt.Println()
}
