package main

import (
	"fmt"

	"github.com/vlasove/materials/fp/chap1/fib"
)

func main() {
	for i := 0; i < 4; i++ {
		fmt.Println("mem:", fib.FibMemoized(i))
		fmt.Println("chan:", fib.FibChanneled(i))
		fmt.Println("simple", fib.FibSimple(i))
	}
}
