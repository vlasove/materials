// Package task18 Что выведет данная программа и почему?
package task18

import "fmt"

// someAction ...
func someAction(v []int8, b int8) {
	v[0] = 100
	fmt.Println("Len:", len(v), "Cap:", cap(v))
	v = append(v, b) // здесь перевыделение памяти идет
	fmt.Println("Len:", len(v), "Cap:", cap(v))
}

// Start ...
func Start() {
	a := make([]int8, 10)
	for i := 0; i < 5; i++ {
		a[i] = int8(i) + 1
	}

	someAction(a, 6)

	fmt.Println(a)

	slice := []int8{1}
	fmt.Println("Cap is:", cap(slice))
	slice = append(slice, 2)
	fmt.Println("Cap is:", cap(slice))

}
