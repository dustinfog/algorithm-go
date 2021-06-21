package main

import (
	"fmt"
	"testing"
)

func TestChangeCoin(t *testing.T) {
	coins := []int{1, 2, 5}
	amount := 11
	ret, num := changeCoin(coins, amount)
	fmt.Printf("%+v, %d", ret, num)
}
