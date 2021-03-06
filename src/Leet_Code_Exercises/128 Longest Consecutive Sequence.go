package main

import "fmt"

func main() {
	nums := []int{100, 4, 200, 1, 3, 2}
	fmt.Println(longestConsecutive(nums))
}
func longestConsecutive(nums []int) int {
	res := 0
	tmp := make(map[int]bool)
	for _, v := range nums {
		tmp[v] = true
	}
	for _, v := range nums {
		if !tmp[v-1] { // v is lower bound of consecutive elements
			temp := 0
			for tmp[v] {
				temp++
				v++
			}
			if temp > res {
				res = temp
			}
		}
	}
	return res
}
