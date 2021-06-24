package main

import "fmt"

func MaxAscSeq(nums []int) []int {
	var longestSeq []int
	for i, num := range nums {
		if i == 0 {
			longestSeq = []int{num}
			fmt.Printf("%v", longestSeq)
		} else {
			if num > longestSeq[len(longestSeq)-1] {
				longestSeq = append(longestSeq, num)
				fmt.Printf("%v", longestSeq)
			} else {
				for i, tmp := range longestSeq {
					if num < tmp {
						longestSeq[i] = num
						fmt.Printf("%v", longestSeq)
					}
				}
			}
		}
	}

	return longestSeq
}
