package main

import "fmt"

func main() {
	prices := []int{7, 1, 5, 3, 6, 4}
	fmt.Println(maxProfit(prices))
}

func maxProfit(prices []int) int {
	l, r := 0, 1
	maxP := 0

	for r < len(prices) {
		if prices[l] < prices[r] {
			profit := prices[r] - prices[l]
			maxP = Max(maxP, profit)
		} else {
			l = r
		}
		r++
	}
	return maxP
}

func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
