package main

import "fmt"

const (
	a = 2017 + iota
	b = 2017 + iota
	c = 2017 + iota
	d = 2017 + iota
)

func main() {
	//Exercise_01()
	//Exercise_02()
	//Exercise_03()
	//Exercise_04()
	//Exercise_05()
	Exercise_06()
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

func Exercise_04() {
	/*	Hands-on exercise #4
			Write a program that
			● assigns an int to a variable
			● prints that int in decimal, binary, and hex
			● shifts the bits of that int over 1 position to the left, and assigns that to a variable
			● prints that variable in decimal, binary, and hex
		solution: https://play.golang.org/p/Ms964T8SbH
		video: 034*/
	var x int = 45
	fmt.Printf("%d\t %b\t %#x \n", x, x, x)
	y := x << 1
	fmt.Printf("%d\t %b\t %#x", y, y, y)
}

func Exercise_05() {
	/*	Hands-on exercise #5
		Create a variable of type string using a raw string literal. Print it.
			solution: https://play.golang.org/p/dLy36A-V-w
	video: 035*/
	a := "Raw string literal"
	fmt.Println(a)
}

func Exercise_06() {
	/*	Hands-on exercise #6
			Using iota, create 4 constants for the NEXT 4 years. Print the constant values.
				solution: https://play.golang.org/p/MDLF3v5EGT
		video: 036
	*/
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(d)
}
