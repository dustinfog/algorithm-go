package main

// https://leetcode-cn.com/problems/cyclically-rotating-a-grid/
//
func rotateGrid(grid [][]int, k int) [][]int {
	g := NewGrid(len(grid[0]), len(grid))
	for r := g.GetRingNum() - 1; r >= 0; r-- {
		ring := g.GetRing(r)
		l := ring.Len()
		k := k % l

		for i := 0; i < l/2; i++ {
			sx, sy := ring.GetCoord(i)
			tx, ty := ring.GetCoord(l - i - 1)

			grid[sy][sx], grid[ty][tx] = grid[ty][tx], grid[sy][sx]
		}

		for i := 0; i < (l-k)/2; i++ {
			sx, sy := ring.GetCoord(i)
			tx, ty := ring.GetCoord(l - k - i - 1)

			grid[sy][sx], grid[ty][tx] = grid[ty][tx], grid[sy][sx]
		}

		for i := 0; i < k/2; i++ {
			sx, sy := ring.GetCoord(l - k + i)
			tx, ty := ring.GetCoord(l - i - 1)

			grid[sy][sx], grid[ty][tx] = grid[ty][tx], grid[sy][sx]
		}
	}

	return grid
}

type GridRing struct {
	sx, sy, w, h int
	l            int
}

func NewGridRing(sx, sy, w, h int) *GridRing {
	return &GridRing{
		sx: sx,
		sy: sy,
		w:  w,
		h:  h,
		l:  2 * (w + h - 2),
	}
}

func (g *GridRing) GetCoord(i int) (int, int) {
	if i < g.w {
		return g.sx + i, g.sy
	}

	if i < g.w+g.h-1 {
		return g.sx + g.w - 1, g.sy + i - g.w + 1
	}

	if i < g.w*2+g.h-2 {
		return g.sx + 2*g.w + g.h - i - 3, g.sy + g.h - 1
	}

	return g.sx, g.sy + 2*g.w + 2*g.h - i - 4
}

func (g *GridRing) Len() int {
	return g.l
}

type Grid struct {
	w int
	h int
}

func NewGrid(w, h int) *Grid {
	return &Grid{
		w: w,
		h: h,
	}
}

func (g *Grid) GetRingNum() int {
	return min(g.w, g.h) / 2
}

func (g *Grid) GetRing(n int) *GridRing {
	return NewGridRing(n, n, g.w-n*2, g.h-n*2)
}

//
//func fmtGrid(grid [][]int) string {
//	var str string
//	for i, h := 0, len(grid); i < h; i++ {
//		for j, w := 0, len(grid[i]); j < w; j++ {
//			str += strconv.Itoa(grid[i][j]) + "\t"
//		}
//		str += "\n"
//	}
//
//	str += "\n"
//	return str
//}
