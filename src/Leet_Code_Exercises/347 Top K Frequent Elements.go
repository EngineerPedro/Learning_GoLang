package main

import "fmt"

func main() {
	nums := []int{1, 1, 1, 2, 2, 3}
	k := 2
	fmt.Println(topKFrequent(nums, k))
}

func topKFrequent(nums []int, k int) []int {

	//make a map the store value (the int as the key) and the value will be the frequecny at which it occurs
	d := make(map[int]int)
	for _, x := range nums { //range through that slice getting the frequency for each integer
		d[x]++
	}

	//make a bucket the length of the array/slice
	b := make([][]int, len(nums)+1)
	for x, f := range d { //go through the maps values using it as the index to store the key as the value
		b[f] = append(b[f], x) // store the key as the value
	}

	// make a slice of size 1 SO THAT it can be incremented to the size k
	res := make([]int, 0)
	for i := len(b) - 1; i >= 0; i-- { // loop through the length of the bucket BACKWARDS to get the proper order
		for _, x := range b[i] { // got through every value in the slice
			if len(res) == k { // check if we are at the desired length of the array/slice to return
				return res
			}
			res = append(res, x) //grab each value and put it into the returning slice
		}
	}
	return res
}
