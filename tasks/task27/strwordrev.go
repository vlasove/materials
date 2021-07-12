// Package task27 Написать программу,
// которая переворачивает
// слова в строке (snow dog sun - sun dog snow).
package task27

import (
	"fmt"
	"strings"
)

// ReverseWords ...
func ReverseWords(s string) string {
	words := strings.Split(s, " ")
	for i, j := 0, len(words)-1; i < j; i, j = i+1, j-1 {
		words[i], words[j] = words[j], words[i]
	}
	return strings.Join(words, " ")
}

// Start ...
func Start() {
	sample := "Витя ехал далеко! Витя ехал за пивком:)"
	revSample := ReverseWords(sample)
	fmt.Println("Origin:", sample)
	fmt.Println("Reversed:", revSample)
}
