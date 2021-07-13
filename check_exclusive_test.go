package main

import "testing"

func TestCheckExclusive(t *testing.T) {
	if !checkExclusive([]int{1, 2}, []int{1, 2}) {
		t.Fatal(1)
	}

	if !checkExclusive([]int{0}, []int{1, 2}) {
		t.Fatal(2)
	}

	if !checkExclusive([]int{1}, []int{1, 2}) {
		t.Fatal(3)
	}

	if checkExclusive([]int{0, 1}, []int{1, 2}) {
		t.Fatal(4)
	}
}
