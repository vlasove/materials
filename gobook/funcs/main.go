package main

import "fmt"

func half(n int) (int, bool) {
	return n / 2, n%2 == 0
}

func max(nums ...int) int {
	if len(nums) == 0 {
		return 0
	}
	m := nums[0]
	for _, v := range nums {
		if v > m {
			m = v
		}
	}
	return m
}

func makeOddGenerator() func() uint {
	i := uint(1)
	return func() uint {
		ret := i
		i += 2
		return ret
	}
}

func fib(n uint) uint {
	if n <= 1 {
		return 1
	}
	return fib(n-1) + fib(n-2)
}

func punisher() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println(r)
		}
	}()
	panic("I'm panic!")
}
func main() {
	defer fmt.Println("1")
	defer fmt.Println("2")
	defer fmt.Println("3")

	fmt.Println(half(2))
	fmt.Println(half(13))
	fmt.Println(max(10, 20, 1, 3, 4, 2))
	odd := makeOddGenerator()
	fmt.Println(odd(), odd(), odd(), odd())
	for i := 0; i < 10; i++ {
		fmt.Printf("%d and fib:%d\n", i, fib(uint(i)))
	}
	punisher()
}
