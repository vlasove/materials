// Package task12 Что выводит
// данная программа и почему?
package task12

import "fmt"

func update(p *int) {
	b := 2
	fmt.Println("Ptr:", p, "Value in funciton:", *p)
	p = &b // local p value of ptr to b
	fmt.Println("Ptr:", p, "Local value:", *p)
}

// Start ...
func Start() {
	a := 1
	p := &a

	fmt.Println(*p) // *p -> 1, p -> ptr, in main
	update(p)
	fmt.Println(*p) // *p -> 1, p -> ptr, in main (not changed)

}
