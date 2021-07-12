// Package task26 Написать программу,
// которая переворачивает строку.
// Символы могут быть unicode.
package task26

import "fmt"

// Reverse ...
func Reverse(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

// Start ...
func Start() {
	sample := "Витя ехал далеко! Витя ехал за пивком:)"
	revSample := Reverse(sample)
	fmt.Println("Origin:", sample)
	fmt.Println("Reversed:", revSample)
}
