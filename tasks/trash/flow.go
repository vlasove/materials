package trash

import (
	"fmt"
	"sort"
)

// Value ...
type Value struct {
	val       int
	something string
}

var (
	values = []Value{
		{val: 10, something: "poop"},
		{-2, "abcd"},
		{3, "qwqerty"},
	}
)

// Start ...
func Start() {
	fmt.Println("Usorted:", values)
	sort.Slice(values, func(i, j int) bool {
		return values[i].val <= values[j].val
	})
	fmt.Println("Sorted:", values)

	valsToSearch := Value{10, "pooper"}
	idx := sort.Search(len(values), func(i int) bool {
		return values[i] == valsToSearch
	})
	fmt.Println("Searched at:", idx)
}
