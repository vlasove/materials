// Package task34 Написать программу,
// которая проверяет, что все символы в строке уникальные.
package task34

import (
	"fmt"
	"unicode/utf8"

	"github.com/vlasove/materials/tasks/task14"
)

// IsUniqueRune ...
func IsUniqueStr(msg string) bool {
	s := task14.New()
	for _, v := range msg {
		s.Add(string(v))
	}
	return s.Size() == utf8.RuneCountInString(msg)
}

// Start ...
func Start() {
	msg := "Ola,it's me!" // "Ola, it's me, Mario!"

	if IsUniqueStr(msg) {
		fmt.Println(msg, "has unique chars")
	} else {
		fmt.Println(msg, "has doubled chars")
	}
}
