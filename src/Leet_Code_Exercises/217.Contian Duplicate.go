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
	hm := make(map[int]int) //Hashmap that'll store our values into the key and also allow us to check that key
	// immediately if the hashmap contains a matching key already

	for _, value := range nums { //takes us through every index of the array giving us the value to be examined
		if _, ok := hm[value]; ok { //using the comma OK we check if their exists a matching key to the value
			//inside of the hash map if it exists the problem is solved and we return rtue
			return true
		} else { //if not then we add that key to the hashmap
			hm[value] = nums[value]
		}
	}
	return false // in the situation we have added ALL values contained in the array / slice into the hashmap
	//and none of the keys have matched  return false
}
