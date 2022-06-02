package main

import "fmt"

/*
Given an array of integers nums and an integer target, return indices of the two numbers such that they add up to target.

You may assume that each input would have exactly one solution, and you may not use the same element twice.

You can return the answer in any order.

What Do you want?
*return indices

*/

//Verify the constraints
/*
2 <= nums.length <= 104

-109 <= nums[i] <= 109

-109 <= target <= 109
*/

//Write out test cases
/*
Example 1:

Input: nums = [2,7,11,15], target = 9
Output: [0,1]
Explanation: Because nums[0] + nums[1] == 9, we return [0, 1].

Example 2:

Input: nums = [3,2,4], target = 6
Output: [1,2]

Example 3:

Input: nums = [3,3], target = 6
Output: [0,1]
*/

//Work out solution w/o code::write out the logic
/*
x     +    y  =   z
index + index = target

translated:
Pointer_1 + Pointer_2 = Target_Value
x         +    y      =   z

Target_value - Pointer_1 = Pointer_2
z            -  x        = y

This will lead us to a solution that is going be just BRUTE FORCE
meaning nested for loops time complexity of O(N^2)

So what about this:
*Target - Pointer_1  will produce the variable "Searching_For"
*Searching_For will allow us to produce ONE item that we are just searching for in
the array helping us full abandon the two pointer method
*Even better create a hashtable so you do not have iterate through every index
this allows us to just check if their a key that contains the value we are searching for
*/

//Pseudocode
/*
-convert the slice into a map
-create a loop to go through each index of the slice or map
-create the search_for  variable by subtracting the current index from the target
*/

func main() {
	target := 25
	nums := []int{6, 8, 9, 10, 34, 60, 15}
	fmt.Println(twoSum(nums, target))

}

func twoSum(nums []int, target int) []int {
	//Create a map
	hm := make(map[int]int)
	//range through all values in the slice
	//addend is each value in the slice nums
	for index01, Current_Value := range nums {
		//Create variable to let us know what we are looking for
		Search_For := target - Current_Value
		//comma ok idiom to make sure at the index/key
		//the value we are looking for exists
		if index02, ok := hm[Search_For]; ok {
			//returns the indexs of the spot in the slices that meet the requirements
			return []int{index01, index02}
		}
		hm[Current_Value] = index01
	}

	return []int{-1, -1}
}
