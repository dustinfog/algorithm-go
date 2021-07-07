package main

import (
	"testing"
)

func TestGridRing(t *testing.T) {

	r := NewGridRing(0, 0, 4, 4)

	if r.Len() != 12 {
		t.Fatal(12)
	}

	x, y := r.GetCoord(0)
	if x != 0 || y != 0 {
		t.Fatal(0)
	}

	x, y = r.GetCoord(4)
	if x != 3 || y != 1 {
		t.Fatal(4)
	}

	x, y = r.GetCoord(8)
	if x != 1 || y != 3 {
		t.Fatal(8)
	}

	r = NewGridRing(1, 1, 2, 2)

	if r.Len() != 4 {
		t.Fatal(4)
	}

	x, y = r.GetCoord(0)
	if x != 1 || y != 1 {
		t.Fatal(0)
	}

	x, y = r.GetCoord(2)
	if x != 2 || y != 2 {
		t.Fatal(2)
	}

	r = NewGridRing(1, 1, 2, 8)
	//01
	//52
	//43
	//34
	//25
	//16
	//07
	//98

	if r.Len() != 16 {
		t.Fatal(r.Len())
	}

	x, y = r.GetCoord(7)
	if x != 2 || y != 7 {
		t.Fatal(x, y)
	}

	x, y = r.GetCoord(13)
	if x != 1 || y != 4 {
		t.Fatal(x, y)
	}
}

func TestRotateGrid(t *testing.T) {
	grid := [][]int{
		{40, 10},
		{30, 20},
	}

	grid = rotateGrid(grid, 1)
	t.Logf("%v", grid)
}

func TestRotateGrid1(t *testing.T) {
	grid := [][]int{
		{1, 2, 3, 4},
		{5, 6, 7, 8},
		{9, 10, 11, 12},
		{13, 14, 15, 16},
	}

	//[[3 5 9 13]
	//[2 11 10 14]
	//[1 7 6 16]
	//[15 4 8 12]]
	grid = rotateGrid(grid, 2)
	t.Logf("%v", grid)
}

func TestRotateGrid2(t *testing.T) {
	grid := [][]int{
		{10, 1, 4, 8},
		{6, 6, 3, 10},
		{7, 4, 7, 10},
		{1, 10, 6, 1},
		{2, 1, 1, 10},
		{3, 8, 9, 2},
		{7, 1, 10, 10},
		{7, 1, 4, 9},
		{2, 2, 4, 2},
		{10, 7, 5, 10},
	}

	//[[3 5 9 13]
	//[2 11 10 14]
	//[1 7 6 16]
	//[15 4 8 12]]
	grid = rotateGrid(grid, 1)
	t.Logf("%v", grid)
}

//
//func TestRotateGrid2(t *testing.T) {
//	grid := [][]int{
//		{7, 8, 8, 10, 4, 7, 2, 8},
//		{4, 8, 9, 9, 1, 5, 9, 4},
//		{8, 9, 4, 8, 9, 9, 1, 5},
//		{1, 2, 7, 8, 10, 1, 1, 10},
//		{7, 6, 6, 8, 2, 5, 5, 4},
//		{7, 7, 3, 3, 6, 10, 10, 6},
//		{2, 5, 2, 9, 3, 9, 10, 2},
//		{7, 5, 8, 1, 3, 3, 4, 7},
//	}
//
//	//[
//	//	[10,4,7,2,8,4,5,10],
//	//	[7,1,5,9,1,1,5,4],
//	//	[8,9,9,1,5,10,10,6],
//	//	[8,8,9,8,10,6,10,2],
//	//	[4,9,8,8,2,3,9,7],
//	//	[8,9,4,7,6,3,3,4],
//	//	[1,2,6,7,5,2,9,3],
//	//	[7,7,2,7,5,8,1,3]
//	//]
//	grid = rotateGrid(grid, 3)
//	s, _ := json.Marshal(grid)
//	t.Logf("%s", s)
//}
