package main

import "fmt"

// FindMin ...
func FindMin(sl []int) int {
	m := sl[0]
	for _, v := range sl {
		if v < m {
			m = v
		}
	}
	return m

}
func main() {
	nums := []int{
		48, 96, 86, 68,
		57, -2, 63, 70,
		37, 34, 83, 27,
		19, 97, 9, 17,
	}

	fmt.Println("Min elem is", FindMin(nums))
}
