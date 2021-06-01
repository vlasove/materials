package helper

import "testing"

// assertCorretMessage used to assert when strings diffrent [got != want]
func assertCorrectMessage(t *testing.T, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
func TestHello(t *testing.T) {
	t.Run("saying hello with name", func(t *testing.T) {
		got := Hello("Bob")
		want := "Hello, Bob!"

		assertCorrectMessage(t, got, want)
	})

	t.Run("saying hello to noboy", func(t *testing.T) {
		got := Hello("")
		want := "Hello, World!"

		assertCorrectMessage(t, got, want)
	})

}
