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

func TestChangeCoin1(t *testing.T) {
	coins := []int{1, 2, 5}
	amount := 11
	ret, num := changeCoin1(coins, amount)
	fmt.Printf("%+v, %d", ret, num)
}

func TestChangeCoin2(t *testing.T) {
	coins := []int{1, 2, 5}
	amount := 5
	ret := changeCoin2(coins, amount)
	fmt.Printf("%+v", ret)
}
