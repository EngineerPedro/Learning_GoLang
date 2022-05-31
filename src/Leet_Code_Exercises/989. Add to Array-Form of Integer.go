package main

import "fmt"

func main() {
	A := []int{2, 5, 5}
	K := 26
	fmt.Println(addToArrayForm(A, K))
}

func addToArrayForm(A []int, K int) []int {
	//get the length of the array passed in and substract by one
	i := len(A) - 1
	//create a slice to store your solution
	res := []int{}
	//loop through that array adding K each time
	for i >= 0 || K > 0 {
		if i >= 0 {
			K += A[i]
		}
		//Add the result to the new
		res = append(res, K%10)
		//Take us to the next 10th place
		K /= 10
		//decrement down the array
		i--
	}
	for l, r := 0, len(res)-1; l < r; l, r = l+1, r-1 {
		res[l], res[r] = res[r], res[l]
	}
	return res
}
