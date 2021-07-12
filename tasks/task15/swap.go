// Package task15 Поменять местами
// два числа без создания временной переменной.
package task15

import "fmt"

// Start ...
func Start() {
	a, b := 10, 20
	fmt.Println("a:", a, "b:", b)
	a, b = b, a
	fmt.Println("a:", a, "b:", b)
}
