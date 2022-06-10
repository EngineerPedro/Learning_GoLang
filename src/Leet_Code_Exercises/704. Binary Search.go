package main

import "fmt"

func main() {
	nums := []int{-1, 0, 3, 5, 9, 12}
	target := 9
	fmt.Println(search(nums, target))
}

func search(nums []int, target int) int {
	middle, left, right := 0, 0, len(nums)-1

	// While the left is less this or equal to our right pointer
	for left <= right {
		// Set to the middle of the array
		middle = left + (right-left)/2

		// Check if the middle equals the target
		if nums[middle] == target {
			// If so return our answer
			return middle
		}
		// Check if the middle is less then target
		if nums[middle] < target {
			// If so move the left pointer to the right
			left = middle + 1
		} else {
			// if not move the right pointer left
			right = middle - 1
		}
	}
	// if the target is never found return -1
	return -1
}
