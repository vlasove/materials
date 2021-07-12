// Package task11 Написать пересечение
// двух неупорядоченных массивов.
package task11

import "fmt"

var (
	firstArr  = [5]int{2, 3, 5, 4, 1}
	secondArr = [3]int{2, 8, 1}
)

// Intersection ...
func Intersection(first, second []int) []int {
	ans := []int{}

	hash := make(map[int]bool)
	for _, elem := range first {
		hash[elem] = true
	}

	for _, elem := range second {
		if hash[elem] {
			ans = append(ans, elem)
		}
	}
	//ans = sanitize(ans)
	return ans
}

// sanitize for delete duplicates
// func sanitize(elements []int) (nodups []int) {
// 	encountered := make(map[int]bool)
// 	for _, element := range elements {
// 		if !encountered[element] {
// 			nodups = append(nodups, element)
// 			encountered[element] = true
// 		}
// 	}
// 	return
// }

// Start ...
func Start() {
	res := Intersection(firstArr[:], secondArr[:])
	fmt.Println("Intersection:", res)
}
