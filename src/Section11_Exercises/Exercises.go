package main

import (
	"fmt"
)

func main() {
	//Exercise_01()
	//Exercise_02()
	Exercise_03()
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

func Exercise_02() {
	/*	Hands-on exercise #2
			● Using a COMPOSITE LITERAL:
			● create a SLICE of TYPE int
			● assign 10 VALUES
			● Range over the slice and print the values out.
			● Using format printing
			○ print out the TYPE of the slice
		solution: https://play.golang.org/p/sAQeFB7DIs
		video: 072*/
	x := []int{43, 49, 55, 32, 15, 57, 69, 35, 12, 10}

	for i, v := range x {
		fmt.Println(i, v)
	}
	fmt.Printf("%T\n", x)

}

func Exercise_03() {
	/*Hands-on exercise #3
	Using the code from the previous example, use SLICING to create the following new slices
	which are then printed:
	● [42 43 44 45 46]
	● [47 48 49 50 51]
	● [44 45 46 47 48]
	● [43 44 45 46 47]
	solution: https://play.golang.org/p/SGfiULXzAB
	video: 073*/

	x := []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51}

	fmt.Println(x[:5])
	fmt.Println(x[5:])
	fmt.Println(x[2:7])
	fmt.Println(x[1:6])

}
