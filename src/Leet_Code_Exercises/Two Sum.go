package main

import "fmt"

func main() {
	target := 25
	nums := []int{6, 8, 9, 10, 34, 60, 15}
	fmt.Println(twoSum(nums, target))

}

func twoSum(nums []int, target int) []int {
	lookup := make(map[int]int)
	for i, v := range nums {
		j, ok := lookup[-v]
		lookup[v-target] = i
		if ok {
			return []int{j, i}
		}
	}
	return []int{}
}
