package main

import (
	"fmt"
)

func main() {
	nums := []int{1, 2, 3, 1}
	fmt.Println(containsDuplicate(nums))
}

///Brute force
/*func containsDuplicate(nums []int) bool {
	for i := 0; i < len(nums); i++ {
		CurrentValue := nums[i]
		for _, num := range nums {
			if CurrentValue == num {
				return true
			}
		}
	}
	return false
}*/

//Using Hashmap Data Structure
func containsDuplicate(nums []int) bool {
	hm := make(map[int]int)
	for _, value := range nums {
		if _, ok := hm[value]; ok {
			return true
		} else {
			hm[value] = nums[value]
		}
	}
	return false
}
