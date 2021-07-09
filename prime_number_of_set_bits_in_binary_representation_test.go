package main

import "testing"

func TestCountPrimeSetBits(t *testing.T) {
	ret := countPrimeSetBits(6, 10)
	if ret != 4 {
		t.Fatal(ret)
	}

	ret = countPrimeSetBits(10, 15)
	if ret != 5 {
		t.Fatal(ret)
	}
}
