package main

import "fmt"

func main() {
	Exercise_01()
}

func Exercise_01() {
	
	/*Hands-on exercise #1
		Print every number from 1 to 10,000
	solution: https://play.golang.org/p/voDiuiDGGw
	video: 050*/

	for x := 0; x <= 100; x++ {
		fmt.Println(x)
	}
}
