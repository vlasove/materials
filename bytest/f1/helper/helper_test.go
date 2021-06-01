package helper

import (
	"fmt"
	"testing"
)

// assertCorretMessage used to assert when strings diffrent [got != want]
func assertCorrectMessage(t *testing.T, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
func TestHello(t *testing.T) {
	t.Run("saying hello with name", func(t *testing.T) {
		got := Hello("Bob", "")
		want := "Hello, Bob!"

		assertCorrectMessage(t, got, want)
	})

	t.Run("saying hello to noboy", func(t *testing.T) {
		got := Hello("", "")
		want := "Hello, World!"

		assertCorrectMessage(t, got, want)
	})

	t.Run("saying hola to somebody on spanish", func(t *testing.T) {
		got := Hello("Chris", "Spanish")
		want := "Hola, Chris!"

		assertCorrectMessage(t, got, want)
	})
	t.Run("saying hola to nobody on spanish", func(t *testing.T) {
		got := Hello("", "Spanish")
		want := "Hola, World!"

		assertCorrectMessage(t, got, want)
	})

	t.Run("saying bonjour to somebody on french", func(t *testing.T) {
		got := Hello("Anna", "French")
		want := "Bonjour, Anna!"

		assertCorrectMessage(t, got, want)
	})
	t.Run("saying bonjour to nobody on french", func(t *testing.T) {
		got := Hello("", "French")
		want := "Bonjour, World!"

		assertCorrectMessage(t, got, want)
	})

}

func TestGreetingPrefix(t *testing.T) {
	t.Run("english test", func(t *testing.T) {
		got := greetingPrefix("English")
		want := englishHelloPrefix
		assertCorrectMessage(t, got, want)
	})
	t.Run("english (default) test", func(t *testing.T) {
		got := greetingPrefix("Russian")
		want := englishHelloPrefix
		assertCorrectMessage(t, got, want)
	})
	t.Run("spanish test", func(t *testing.T) {
		got := greetingPrefix("Spanish")
		want := spanishHelloPrefix
		assertCorrectMessage(t, got, want)
	})
	t.Run("french test", func(t *testing.T) {
		got := greetingPrefix("French")
		want := frenchHelloPrefix
		assertCorrectMessage(t, got, want)
	})

}

func ExampleHello() {
	msg := Hello("Bob", "Spanish")
	fmt.Println(msg)
	// Output: Hola, Bob!
}
