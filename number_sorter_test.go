package main

import "testing"

func TestSortNumber(t *testing.T) {
	given := []int{9, 8, 6, 1, 2, 5, 4, 3, 7}
	expected := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	sortNumber(given)

	for idx := range given {
		if given[idx] != expected[idx] {
			t.Fatalf("got %d, want %d", given[idx], expected[idx])
		}
	}
}
