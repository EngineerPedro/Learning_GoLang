package main

import "sort"

func main() {

}

func lastStoneWeight(stones []int) int {

	for len(stones) > 1 {
		sort.Ints(stones)
		x := stones[len(stones)-2]
		y := stones[len(stones)-1]
		stones = append(stones[:len(stones)-2])
		if x != y {
			y = abs(y - x)
			stones = append(stones, y)
		}
	}
	if len(stones) > 0 {
		return stones[0]
	}
	return 0
}
func abs_01(n int) int {
	if n < 0 {
		return n * -1
	} else {
		return n
	}
}
