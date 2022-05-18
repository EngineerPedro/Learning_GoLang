package main

import "fmt"

func main() {
	Exercise_01()
}

func Exercise_01() {
	/*
		Hands-on exercise #1
		● Create a value and assign it to a variable.
		● Print the address of that value.
			code: https://play.golang.org/p/57kW8xd0qT
		video: 116
	*/
	x := 42
	fmt.Println(x)
	fmt.Println(&x)
}
