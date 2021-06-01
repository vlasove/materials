package main

import "fmt"

func main() {
	for i := 0; i <= 100; i++ {
		switch {
		case i%3 == 0 && i%5 == 0:
			fmt.Printf("%s ", "FizzBuzz")
		case i%3 == 0:
			fmt.Printf("%s ", "Fizz")
		case i%5 == 0:
			fmt.Printf("%s ", "Buzz")
		default:
			fmt.Printf("%d ", i)
		}
	}
	fmt.Println()
}
