package main

//https://leetcode-cn.com/problems/prime-number-of-set-bits-in-binary-representation/

const PrimMask = 1<<2 |
	1<<3 |
	1<<5 |
	1<<7 |
	1<<11 |
	1<<13 |
	1<<17 |
	1<<19

func countPrimeSetBits(left int, right int) int {
	ret := 0
	for i := left; i <= right; i++ {
		if PrimMask&(1<<getPrimeSetBits(i)) != 0 {
			ret++
		}
	}

	return ret
}

func getPrimeSetBits(num int) int {
	ret := 0
	for num != 0 {
		num = num & (num - 1)
		ret++
	}

	return ret
}
