package main

import (
	"fmt"
)

func main() {
	Exercise_01()
}

func Exercise_01() {

	/*	Hands-on exercise #1
			● Using a COMPOSITE LITERAL:
			● create an ARRAY which holds 5 VALUES of TYPE int
			● assign VALUES to each index position.
			● Range over the array and print the values out.
			● Using format printing
			○ print out the TYPE of the array
		solution: https://play.golang.org/p/tD0SzV3hdf
		video: 071*/

	var x [5]int
	y := 1
	for i := 0; i <= 4; i++ {
		x[i] = y
		y++
	}

	fmt.Println(x)

	for i, v := range x {
		fmt.Println(i, v)
	}

	fmt.Printf("%T\n", x)
}
