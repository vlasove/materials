package helper

import "fmt"

const (
	englishHelloPrefix = "Hello"
	spanishHelloPrefix = "Hola"
	frenchHelloPrefix  = "Bonjour"
	spanish            = "Spanish"
	french             = "French"
)

// greetingPrefix function choosing right prefix by language
// default it english
func greetingPrefix(lanuage string) (prefix string) {
	switch lanuage {
	case spanish:
		prefix = spanishHelloPrefix
	case french:
		prefix = frenchHelloPrefix
	default:
		prefix = englishHelloPrefix
	}
	return
}

// Hello function ...
func Hello(name string, language string) string {
	if name == "" {
		name = "World"
	}
	return fmt.Sprintf("%s, %s!", greetingPrefix(language), name)
}
