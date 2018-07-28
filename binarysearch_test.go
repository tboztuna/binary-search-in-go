package binarysearch

import (
	"testing"
)

var want = 7
var newSlice []int

var tables = []struct {
	a int
	b int
	c int
	d int
	e int
	f int
	g int
	h int
	j int
	k int
}{
	{1, 4, 5, 6, 6, 7, 8, 32, 55, 66},
	{2, 2, 3, 4, 5, 6, 6, 32, 15, 15},
	{10, 12, 12, 20, 20, 21, 22, 32, 1667, 2222},
}

func TestRecursiveBinarySearch(t *testing.T) {

	for _, table := range tables {
		newSlice = []int{
			table.a,
			table.b,
			table.c,
			table.d,
			table.e,
			table.f,
			table.g,
			table.h,
			table.j,
			table.k,
		}
		got := recursiveBinarySearch(newSlice, 32, 0, len(newSlice)-1)

		if got != want {
			t.Errorf("An error occurred while searching, got: %d, want: %d.", got, want)
		}
	}
}

func TestIterativeBinarySearch(t *testing.T) {

	for _, table := range tables {
		newSlice = []int{
			table.a,
			table.b,
			table.c,
			table.d,
			table.e,
			table.f,
			table.g,
			table.h,
			table.j,
			table.k,
		}
		got := iterativeBinarySearch(newSlice, 32, 0, len(newSlice)-1)

		if got != want {
			t.Errorf("An error occurred while searching, got: %d, want: %d.", got, want)
		}
	}
}
