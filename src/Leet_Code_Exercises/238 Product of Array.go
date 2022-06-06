package main

import "fmt"

func main() {
	nums := []int{1, 2, 3, 4}
	fmt.Println(productExceptSelf(nums))
}

func productExceptSelf(nums []int) []int {
	//create a slice of ints that'll store your result
	res := make([]int, len(nums))
	prefix := 1
	for i := 0; i < len(nums); i++ { //loop through the length of the passed in nums array front to back
		res[i] = prefix //set each value of the current index equal to one
		prefix *= nums[i]
	}
	postfix := 1
	for i := len(nums) - 1; i >= 0; i-- { //loop through the length of the passed in nums array back to front
		res[i] *= postfix
		postfix *= nums[i]
	}
	return res
}
