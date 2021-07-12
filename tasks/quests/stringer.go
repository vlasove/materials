// Package quests
// Какой самый эффективный способ работы с объединением строк?
package quests

import (
	"fmt"
	"strings"
	"unsafe"
)

type D interface{}

func Swap() {}

// Start ...
func Start() {
	var b strings.Builder
	fmt.Printf("Initial string: %q\n", b.String())
	for i := 3; i >= 1; i-- {
		b.WriteString(fmt.Sprintf("%d...", i))
		//fmt.Fprintf(&b, "%d...", i)
	}
	b.WriteString("ignition")
	fmt.Printf("Appended string: %q\n", b.String())
	var v interface{}
	fmt.Printf("Size of empty interface: %v\n", unsafe.Sizeof(v))
	fmt.Printf("Type of %T\n", Swap)

	m := make(map[int]int)
	m[0] = 1
	m[1] = 20
	m[2] = 30
	fmt.Println(m)
	for k, v := range m {
		fmt.Println(k, v)
	}

	// q := new([100]int)
	// fmt.Println(q[:50])
	// qq := make([]int, 100)[:50]
	// fmt.Println(qq)

	// q := new(map[int]int)
	// (*q)[1] = 20
	// fmt.Println(q)

	// q := make(chan int, 2)
	// fmt.Printf("Val %v Type %T\n", q, q)
	// go func() {
	// 	fmt.Println("In goroutine")
	// 	q <- 1
	// }()
	// <-q
	// time.Sleep(1 * time.Second)

}
