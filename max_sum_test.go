package main

import (
	"testing"
)

func TestMaxSum(t *testing.T) {
	max := maxSum([]int{-2, 1, -3, 4, -1, 2, 1, -5, 4})
	if max != 6 {
		t.Fail()
	}
}
