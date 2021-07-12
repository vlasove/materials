// Package task24 Создать слайс с предварительно выделенными 100 элементами.

package task24

import "fmt"

func Start() {
	s := make([]interface{}, 0, 100)
	fmt.Printf("Val: %v Length:%d Allocated: %d\n", s, len(s), cap(s))
}
