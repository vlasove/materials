// Package task30 Удалить i-ый элемент из слайса.
package task30

import "fmt"

var (
	id = 2
)

func Start() {
	slice := []string{"a", "b", "c", "d", "e"}
	slice = append(slice[:id], slice[id+1:]...)
	fmt.Println(slice)
}
