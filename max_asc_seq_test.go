package main

import (
	"fmt"
	"testing"
)

func TestMaxAscSeq(t *testing.T) {
	ret := MaxAscSeq([]int{10, 9, 2, 5, 3, 7, 101, 18})
	if len(ret) != 4 {
		t.Fail()
	}

	fmt.Printf("%v", ret)
}
