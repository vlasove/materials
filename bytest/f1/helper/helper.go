package helper

import "fmt"

const englishHelloPrefix = "Hello"

// Hello function ...
func Hello(name string) string {
	if name == "" {
		name = "World"
	}
	return fmt.Sprintf("%s, %s!", englishHelloPrefix, name)
}
