package main

import "math"

//https://leetcode-cn.com/problems/minimum-skips-to-arrive-at-meeting-on-time/submissions/
// 将时间转化成距离，最终做距离的比较
func minSkips(dist []int, speed, hoursBefore int) (ans int) {
	ceil := func(d int) int {
		return (d + speed - 1) / speed * speed
	}

	n := len(dist)
	dp := make([]int, n)
	dp[0] = 0

	bigOne := math.MaxInt64 >> 1
	for i := 1; i < n; i++ {
		dp[i] = bigOne
	}

	for i, d := range dist {
		for j := min(n-1, i+1); j > 0; j-- {
			// 跳过n步的最小距离 = min(跳了n步这步不跳过的最小距离, 跳了n-1步、这步跳过的最小距离)
			dp[j] = min(ceil(dp[j]+d), dp[j-1]+d)
		}

		// 不跳过，每步都休息
		dp[0] += ceil(d)
	}

	maxD := hoursBefore * speed
	for i, d := range dp {
		// d小意味着时间短
		if d <= maxD {
			return i
		}
	}

	return -1
}

//		remainder:   0,
//	}}
//
//	for i, l := 0, len(dist); i < l; i++ {
//		step := dist[i]
//		newWays := make([]*Way, i+2)
//		for skip, way := range ways {
//			if way == nil {
//				continue
//			}
//
//			currStep := step + way.remainder
//			hours := currStep / speed
//			remainder := currStep % speed
//			hoursBefore = way.hoursBefore - hours
//
//			if hoursBefore < 0 {
//				continue
//			}
//
//			if remainder == 0 {
//				lessOrExchange(newWays, skip, &Way{
//					hoursBefore: hoursBefore,
//					remainder:   0,
//					//skips:       way.skips,
//				})
//			} else {
//				//跳过
//				if i < l-1 {
//					lessOrExchange(newWays, skip+1, &Way{
//						hoursBefore: hoursBefore,
//						remainder:   remainder,
//						//skips:       append(way.skips, i),
//					})
//				}
//
//				// 不跳
//				hoursBefore--
//				if hoursBefore < 0 {
//					continue
//				}
//
//				lessOrExchange(newWays, skip, &Way{
//					hoursBefore: hoursBefore,
//					remainder:   0,
//					//skips:       way.skips,
//				})
//			}
//		}
//		ways = newWays
//	}
//
//	ret := -1
//	for skip, way := range ways {
//		if way != nil && (ret == -1 || skip < ret) {
//			ret = skip
//		}
//	}
//
//	return ret
//}
//
//func lessOrExchange(ways []*Way, skipNum int, way *Way) {
//	if ways[skipNum] == nil || ways[skipNum].hoursBefore < way.hoursBefore || (ways[skipNum].hoursBefore == way.hoursBefore && ways[skipNum].remainder > way.remainder) {
//		ways[skipNum] = way
//	}
//}
