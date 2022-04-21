package main

import "fmt"

func main() {
	Exercise_01()
}

func Exercise_01() {
	/*	Hands-on exercise #2
		Using the following operators, write expressions and assign their values to variables:
		g. ==
		h. <=
		i. >=
		j. !=
		k. <
		l. >
		Now print each of the variables.
			solution: https://play.golang.org/p/76R-poSzaY
	video: 032*/

	x := 40
	y := 50

	fmt.Println(x == y)
	fmt.Println(x <= y)
	fmt.Println(x >= y)
	fmt.Println(x != y)
	fmt.Println(x < y)
	fmt.Println(x > y)
}
