package main

import "testing"

func TestFindRadius(t *testing.T) {
	ret := 0

	//ret = findRadius([]int{1, 2, 3}, []int{2})
	//if ret != 1 {
	//	t.Fatal(ret)
	//}
	//
	//ret = findRadius([]int{1, 2, 3, 4, 8}, []int{1, 4, 9, 10})
	//if ret != 1 {
	//	t.Fatal(ret)
	//}
	//
	//ret = findRadius([]int{1, 5}, []int{2})
	//if ret != 3 {
	//	t.Fatal(ret)
	//}

	ret = findRadius(
		[]int{474833169, 264817709, 998097157, 817129560},
		[]int{197493099, 404280278, 893351816, 505795335},
	)
	if ret != 104745341 {
		t.Fatal(ret)
	}
}
