package main

import "fmt"

func square(x *float64) {
	*x = *x * *x
}
func swap(a, b *int) {
	t := *a
	*a = *b
	*b = t
}
func main() {
	intPtr := new(bool)
	fmt.Println("Addr:", intPtr, "Value:", *intPtr)
	val := 1.5
	square(&val)
	fmt.Println(val)

	var (
		x = 10
		y = 20
	)
	swap(&x, &y)
	fmt.Println(x, y)
}
