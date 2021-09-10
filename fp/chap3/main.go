package main

import (
	"fmt"
	"strings"
)

type StrFunc func(string) string

func Compose(f, g StrFunc) StrFunc {
	return func(s string) string {
		return g(f(s))
	}
}

func main() {
	var recognize = func(name string) string {
		return fmt.Sprintf("Hey ,%s!", name)
	}

	var emphasize = func(line string) string {
		return strings.ToUpper(line)
	}

	var greetGoF = Compose(recognize, emphasize)
	fmt.Println(greetGoF("vAsya"))
}
