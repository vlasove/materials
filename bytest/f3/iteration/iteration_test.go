package iteration

import (
	"fmt"
	"strings"
	"testing"
)

func assertString(t *testing.T, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}

func TestRepeat(t *testing.T) {
	t.Run("repeat 'a' 5 times", func(t *testing.T) {
		got := Repeat("a", 5)
		want := "aaaaa"
		assertString(t, got, want)
	})

	t.Run("repeat 'c' 3 times", func(t *testing.T) {
		got := Repeat("c", 3)
		want := "ccc"
		assertString(t, got, want)
	})
}

func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeat("a", 5)
	}
}

func BenchmarkRepeatStd(b *testing.B) {
	for i := 0; i < b.N; i++ {
		strings.Repeat("a", 5)
	}
}

func ExampleRepeat() {
	res := Repeat("a", 3)
	fmt.Println(res)
	// Output: aaa
}
