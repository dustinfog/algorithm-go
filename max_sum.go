package main

func maxSum(nums []int) int {
	max := 0
	curr := 0
	for i, l := 0, len(nums); i < l; i++ {
		curr += nums[i]

		if curr > max {
			max = curr
		} else if curr < 0 {
			curr = 0
		}
	}

	return max
}
