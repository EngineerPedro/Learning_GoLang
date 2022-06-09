package main

import "fmt"

func main() {
	prices := []int{7, 1, 5, 3, 6, 4}
	fmt.Println(maxProfit(prices))
}

func maxProfit(prices []int) int {
	min, profit := prices[0], 0
	for i := 1; i < len(prices); i++ {
		min = Min(min, prices[i])
		profit = Max(profit, prices[i]-min)
	}
	return profit
}

func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
func Min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
