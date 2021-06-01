package integers

import (
	"fmt"
	"testing"
)

func assertIntEqual(t *testing.T, got, want int) {
	t.Helper()
	if got != want {
		t.Errorf("got '%d' want '%d'", got, want)
	}
}
func TestAdd(t *testing.T) {
	for i := 0; i < 100; i++ {
		got := Add(i, i)
		want := i + i

		assertIntEqual(t, got, want)
	}

}

func ExampleAdd() {
	res := Add(2, 3)
	fmt.Println(res)
	// Output: 5
}
