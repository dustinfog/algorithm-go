package main

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
