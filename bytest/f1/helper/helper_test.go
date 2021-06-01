package helper

import "testing"

func TestHello(t *testing.T) {
	got := Hello("Bob")
	want := "Hello, Bob!"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
