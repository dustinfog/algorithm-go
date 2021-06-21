package main

import "math"

func min(nums ...int) int {
	m := math.MaxInt64
	for _, num := range nums {
		if num < m {
			m = num
		}
	}

	return m
}

func max(nums ...int) int {
	m := math.MinInt64
	for _, num := range nums {
		if num > m {
			m = num
		}
	}

	return m
}
