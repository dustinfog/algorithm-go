package main

import "testing"

func TestReachNumber(t *testing.T) {
	ret := 0

	ret = reachNumber(3)
	if ret != 2 {
		t.Fatal(ret)
	}

	ret = reachNumber(2)
	if ret != 3 {
		t.Fatal(ret)
	}
}

func TestSearchSqrtCeil(t *testing.T) {
	ret := 0

	ret = searchSumSequentialCeil(1)
	if ret != 1 {
		t.Fatal(ret)
	}
	ret = searchSumSequentialCeil(2)
	if ret != 2 {
		t.Fatal(ret)
	}
	ret = searchSumSequentialCeil(3)
	if ret != 2 {
		t.Fatal(ret)
	}
	ret = searchSumSequentialCeil(4)
	if ret != 3 {
		t.Fatal(ret)
	}
	ret = searchSumSequentialCeil(5)
	if ret != 3 {
		t.Fatal(ret)
	}
	ret = searchSumSequentialCeil(6)
	if ret != 3 {
		t.Fatal(ret)
	}
	ret = searchSumSequentialCeil(7)
	if ret != 4 {
		t.Fatal(ret)
	}
	ret = searchSumSequentialCeil(8)
	if ret != 4 {
		t.Fatal(ret)
	}
	ret = searchSumSequentialCeil(9)
	if ret != 4 {
		t.Fatal(ret)
	}
	ret = searchSumSequentialCeil(10)
	if ret != 4 {
		t.Fatal(ret)
	}
	ret = searchSumSequentialCeil(11)
	if ret != 5 {
		t.Fatal(ret)
	}
}
