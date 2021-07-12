package taskx

import "fmt"

func XOR(left, right []int) []int {
	res := []int{}

	hash := make(map[int]bool)
	for _, elem := range right {
		hash[elem] = true
	}

	for _, elem := range left {
		if !hash[elem] {
			res = append(res, elem)
		}
	}
	return res
}

func Start() {
	left := []int{1, 2, 3, 4, 5, 3, 4, 2}
	right := []int{2, 3, 4}
	res := XOR(left, right)
	fmt.Println(res)
}
