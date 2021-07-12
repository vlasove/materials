// Package task19 К каким негативным
// последствиям может привести данный кусок кода и как это исправить?
package task19

import (
	"fmt"
	"strings"
)

var justString string

// createHugeString ...
func createHugeString(size int) string {
	v := new(strings.Builder)
	for i := 0; i < size; i++ {

		v.WriteString("Певасик")
	}
	return v.String()
}

// someFunc ...
func someFunc() {
	v := createHugeString(1 << 9)
	justString = string([]rune(v)[:10])
}

// Start ...
func Start() {
	someFunc()
	fmt.Println(justString)
}
