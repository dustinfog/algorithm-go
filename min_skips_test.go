package main

import (
	"testing"
)

func TestMinSkips(t *testing.T) {
	ret := 0
	ret = minSkips([]int{1, 3, 2}, 4, 2)
	if ret != 1 {
		t.Fatal(ret)
	}

	ret = minSkips([]int{7, 3, 5, 5}, 2, 10)
	if ret != 2 {
		t.Fatal(ret)
	}

	ret = minSkips([]int{7, 3, 5, 5}, 1, 10)
	if ret != -1 {
		t.Fatal(ret)
	}

	ret = minSkips([]int{2, 1, 5, 4, 4, 3, 2, 9, 2, 10}, 6, 7)
	if ret != 7 {
		t.Fatal(ret)
	}

	ret = minSkips([]int{3, 4, 9, 9, 5, 10, 7}, 6, 9)
	if ret != 2 {
		t.Fatal(ret)
	}
}
