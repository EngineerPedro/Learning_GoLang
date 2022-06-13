package main

import (
	"fmt"
)

func main() {
	//Designate how Long the slice will be with user input
	var Length_of_Array int
	fmt.Scanln(&Length_of_Array)

	//Slice that stores values input by user
	numbers := make([]int, Length_of_Array)

	for i := 0; i < Length_of_Array; i++ {
		var newValue int
		fmt.Scanln(&newValue)
		numbers[i] = newValue
	}
	fmt.Println(max_pairwise_product(numbers))
}

func max_pairwise_product(nums []int) int {
	max1 := 0
	max2 := 0

	for i := range nums {
		if nums[i] >= max1 {
			max2 = max1
			max1 = nums[i]
			continue
		}
		if nums[i] > max2 {
			max2 = nums[i]
		}
	}

	return (max1) * (max2)
}
