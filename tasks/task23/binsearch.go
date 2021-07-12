// Package task23
// Написать бинарный поиск встроенными методами языка.

package task23

import (
	"fmt"
	"sort"

	"github.com/vlasove/materials/tasks/task22"
)

// BinSearch или sort.Search
func BinSearch(s []int, searched int) (int, bool) {
	low, high := 0, len(s)-1
	for low <= high {
		based := (low + high) / 2
		switch {
		case s[based] == searched:
			return based, true
		case s[based] < searched:
			low = based + 1
		case s[based] > searched:
			high = based - 1
		}
	}
	return -1, false
}

// Start ...
func Start() {
	s := []int{-100, 121, 20, 30, 4, 5, -7, 8, 900, 2, 3, 5, 10, 12, -121, 121}
	task22.Quicksort(s)
	fmt.Println("Slice to search in:", s)
	elem := 20

	id, ok := BinSearch(s, elem)
	if ok {
		fmt.Println("Searched elem is:", elem, "on position:", id)
	} else {
		fmt.Println("Can not find elem in slice :", elem)
	}
	// Слишком тупо но там тоже бинарный поиск
	fmt.Println("Std search:", sort.SearchInts(s, 20))
}
