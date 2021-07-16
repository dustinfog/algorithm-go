package main

import (
	"sort"
)

func findRadius(houses []int, heaters []int) int {
	sort.Ints(houses)
	sort.Ints(heaters)

	heaterIndex := 0
	maxHeaterIndex := len(heaters) - 1
	houseMaxRadius := 0

	for _, house := range houses {
		var radius int
		heaterIndex, radius = findHouseMinRadius(house, heaters, heaterIndex, maxHeaterIndex)
		if radius < 0 {
			radius = -radius
		}

		if radius > houseMaxRadius {
			houseMaxRadius = radius
		}
	}

	return houseMaxRadius
}

func findHouseMinRadius(house int, heaters []int, start int, end int) (int, int) {
	vs := heaters[start] - house
	ve := heaters[end] - house
	for end > start {
		if vs >= 0 {
			return start, vs
		}

		if ve <= 0 {
			return end, ve
		}

		if end == start+1 {
			if -vs > ve {
				return end, ve
			}

			return start, vs
		}

		mid := (end + start) / 2
		v := heaters[mid] - house
		if v > 0 {
			end = mid
			ve = v
		} else {
			start = mid
			vs = v
		}
	}

	return start, vs
}
