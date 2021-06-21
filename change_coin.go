package main

import (
	"fmt"
)

//数组中给出面额，amount为要兑换到的钱，求兑换出的最少钱币的张数。
//我们采用自下而上的方式进行思考。在计算 F(i) 之前，我们先计算出 　　　F(0)-F(i-1) 的答案。
//例如：coins = [1, 2, 5], amount = 11
func changeCoin(coins []int, amount int) ([]int, int) {
	min := -1
	//var minI = -1
	var minSub []int
	var minCoin int
	for _, coin := range coins {
		if coin > amount {
			continue
		}

		if coin == amount {
			//fmt.Printf("%d", coin)
			return []int{coin}, 1
		}

		sub, r := changeCoin(coins, amount-coin)
		if r < 0 {
			return nil, r
		}

		curr := 1 + r
		if min < 0 || curr < min {
			min = curr
			minSub = sub
			minCoin = coin
		}
	}

	//fmt.Printf("%d", coins[minI])
	return append(minSub, minCoin), min
}

func changeCoin1(coins []int, amount int) ([]int, int) {
	type dpStruct struct {
		value int
		seq   []int
	}
	dp := make([]dpStruct, amount+1)
	dp[0] = dpStruct{
		value: -1,
	}

	for i := 1; i <= amount; i++ {
		dp[i].value = -1
		minCoin := -1
		for _, coin := range coins {
			if coin == i {
				dp[i].value = 1
				dp[i].seq = []int{i}
				minCoin = -1
			} else if coin < i {
				nested := dp[i-coin]
				if nested.value == -1 {
					continue
				}

				curr := nested.value + 1
				if dp[i].value == -1 || curr < dp[i].value {
					dp[i].value = curr
					minCoin = coin
				}
			}
		}

		if minCoin != -1 {
			dp[i].seq = append(dp[i-minCoin].seq, minCoin)
		}

		fmt.Printf("%d: %d\n", i, dp[i])
	}

	return dp[amount].seq, dp[amount].value
}

//给定不同面额的硬币和一个总金额。写出函数来计算可以凑成总金额的硬币组合数。假设每一种面额的硬币有无限个。
//输入: amount = 5, coins = [1, 2, 5]
//输出: 4
//解释: 有四种方式可以凑成总金额:
//5=5
//5=2+2+1
//5=2+1+1+1
//5=1+1+1+1+1
func changeCoin2(coins []int, amount int) [][]int {
	dp := make([][][]int, amount+1)
	for _, coin := range coins {
		for i := 1; i <= amount; i++ {
			if i == coin {
				dp[i] = append(dp[i], []int{coin})
			} else if i > coin {
				prev := dp[i-coin]
				for _, prevSub := range prev {
					dp[i] = append(dp[i], append(prevSub, coin))
				}
			}
		}
	}

	return dp[amount]
}
