// Package task20 ...
package task20

import "fmt"

// Start ...
func Start() {
	slice := []string{"a", "a"}
	func(slice []string) {
		slice = append(slice, "a") // Перевыделние памяти
		slice[0] = "b"             // В локальный slice  b
		slice[1] = "b"             // В локальный slice b
		fmt.Println(slice)         // Выведем локальный slice
	}(slice)
	fmt.Println(slice)

}
