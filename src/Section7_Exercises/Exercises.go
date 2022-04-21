package main

import "fmt"

const (
	a     = 42
	b int = 43
)

func main() {
	//Exercise_01()
	//Exercise_02()
	Exercise_03()
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

func Exercise_02() {

	/*	Hands-on exercise #1
			Write a program that prints a number in decimal, binary, and hex
		solution: https://play.golang.org/p/bAQxcEuK8O
		video: 031*/

	a := 40
	fmt.Println(a)
	fmt.Printf("%d\t %b\t %#x", a, a, a)

}

func Exercise_03() {
	/*	Hands-on exercise #3
		Create TYPED and UNTYPED constants. Print the values of the constants.
			solution: https://play.golang.org/p/NutvJXWUx2
	video: 033*/

	fmt.Println(a, b)
}
