package main

//https://leetcode-cn.com/problems/reach-a-number/solution/dao-da-zhong-dian-shu-zi-by-leetcode/
func reachNumber(target int) int {
	if target < 0 {
		target = -target
	}

	k := searchSumSequentialCeil(target)
	target = sumSequential(k) - target

	if target%2 == 0 {
		// 把其中的一个减法变成加法，并没有多出步骤
		return k
	}

	// target是奇数，需要继续减，使其变成偶数
	// - k 是偶数，则 k + 1奇数，则target - k - 1 是偶数
	// - k 是奇数，则 k + 1偶数，则target - k - 1 是奇数，还需要再走一步
	return k + 1 + k%2
}

func searchSumSequentialCeil(num int) int {
	start, end := 1, num
	for {
		if end <= start {
			return start
		}
		m := (start + end) / 2
		accum := sumSequential(m)
		if accum == num {
			return m
		}

		if accum < num {
			start = m + 1
		} else {
			end = m
		}
	}
}

func sumSequential(k int) int {
	return (1 + k) * k / 2
}
