package main

import "fmt"

func main() {

	fmt.Println("Enter the number () that you want to " +
		"use as an example for recursion :")
	var user_input int

	fmt.Scanln(&user_input)
	fmt.Println(fact(user_input))

}

//Functions

func fact(n int) int {
	//Base case to get out of function
	if n == 0 {
		return 1
	}
	//Recall (functtion calling itself
	return n * fact(n-1)
}
