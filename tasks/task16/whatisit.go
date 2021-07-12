// Package task16 Что выведет программа данная программа?
package task16

import "fmt"

// Start ...
func Start() {
	n := 0
	if true {
		fmt.Println("Addr before:", &n)
		n := 1 // here will be create local var n
		fmt.Println("Addr after:", &n)
		n++
	}
	fmt.Println(n)

}
