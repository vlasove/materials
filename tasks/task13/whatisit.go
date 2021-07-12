// Package task13 Чем завершится данная программа?
package task13

import (
	"fmt"
	"sync"
)

// Start ...
func Start() {
	wg := sync.WaitGroup{}
	for i := 0; i < 5; i++ {
		wg.Add(1)
		// Здесь пропихивается копия wg
		go func(wg sync.WaitGroup, i int) {
			fmt.Println(i)
			wg.Done() // копия done
		}(wg, i)
	}
	wg.Wait() // исходное значение ждет 5 done
	fmt.Println("exit")
}
