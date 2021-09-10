package fib

type Memoized func(int) int

var fibMem = Memoize(fib)

func Memoize(mf Memoized) Memoized {
	//log.Println("call Memoize")
	cache := make(map[int]int)
	return func(key int) int {
		//log.Println("current cache:", cache)
		if val, found := cache[key]; found {
			return val
		}
		temp := mf(key)
		cache[key] = temp
		return temp
	}
}

func fib(x int) int {
	//log.Println("call fib with x :", x)
	if x == 0 {
		return 0
	}
	if x <= 2 {
		return 1
	}
	return fib(x-2) + fib(x-1)
}

func FibMemoized(n int) int {
	//log.Println("call FibMemozied with n:", n)
	return fibMem(n)
}
