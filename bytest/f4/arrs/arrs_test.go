package arrs

import (
	"reflect"
	"testing"
)

func assertIntEqual(t *testing.T, got, want int) {
	t.Helper()

	if got != want {
		t.Errorf("got %d want %d", got, want)
	}
}

func TestSum(t *testing.T) {
	t.Run("collection of 5 size", func(t *testing.T) {
		got := Sum([]int{1, 2, 3, 4, 5})
		want := 15
		assertIntEqual(t, got, want)
	})
	t.Run("collection of any size", func(t *testing.T) {
		got := Sum([]int{5, 5, 5})
		want := 15
		assertIntEqual(t, got, want)
	})
	t.Run("collection of 0 size", func(t *testing.T) {
		got := Sum([]int{})
		want := 0
		assertIntEqual(t, got, want)
	})

}

func assertSliceInt(t *testing.T, got, want []int) {
	t.Helper()
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v want %v", got, want)
	}

}

func TestSumAll(t *testing.T) {
	t.Run("sum all with 2 slices", func(t *testing.T) {
		got := SumAll([]int{1, 2}, []int{0, 9})
		want := []int{3, 9}
		assertSliceInt(t, got, want)
	})
	t.Run("sum all with 1 slice", func(t *testing.T) {
		got := SumAll([]int{1, 2})
		want := []int{3}
		assertSliceInt(t, got, want)
	})
	t.Run("sum all with 0 slices", func(t *testing.T) {
		got := SumAll()
		want := []int{}
		assertSliceInt(t, got, want)
	})

}

func TestSumAllTails(t *testing.T) {
	t.Run("sum all tails 2 slices", func(t *testing.T) {
		got := SumAllTails([]int{1, 2, 3}, []int{3, 4, 5})
		want := []int{5, 9}
		assertSliceInt(t, got, want)
	})
	t.Run("sum all tails 1 slice", func(t *testing.T) {
		got := SumAllTails([]int{1, 2, 3})
		want := []int{5}
		assertSliceInt(t, got, want)
	})
	t.Run("sum all tails 0 slices", func(t *testing.T) {
		got := SumAllTails()
		want := []int{}
		assertSliceInt(t, got, want)
	})
	t.Run("sum all tails empty slice", func(t *testing.T) {
		got := SumAllTails([]int{1}, []int{})
		want := []int{0, 0}
		assertSliceInt(t, got, want)
	})
}
