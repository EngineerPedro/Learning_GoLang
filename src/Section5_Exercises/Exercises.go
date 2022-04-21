package main

import (
	"fmt"
)

var x int
var y string
var z bool

func main() {

	//Exercise_01()
	//Exercise_02()
}

func Exercise_01() {
	/*Hands-on exercise #1
	1. Using the short declaration operator, ASSIGN these VALUES to VARIABLES with the
	IDENTIFIERS “x” and “y” and “z”
	a. 42
	b. “James Bond”
	c. true
	2. Now print the values stored in those variables using
	a. a single print statement
	b. multiple print statements
	code: here’s the solution: https://play.golang.org/p/yYXnWXIQNa
	video: 017 */

	x := 42
	y := "James Bond"
	z := true

	fmt.Println(x)
	fmt.Println(y)
	fmt.Println(z)
}

func Exercise_02() {
	/*	Hands-on exercise #2
			1. Use var to DECLARE three VARIABLES. The variables should have package level
			scope. Do not assign VALUES to the variables. Use the following IDENTIFIERS for the
			variables and make sure the variables are of the following TYPE (meaning they can
			store VALUES of that TYPE).
			a. identifier “x” type int
			b. identifier “y” type string
			c. identifier “z” type bool
			2. in func main
			a. print out the values for each identifier
			b. The compiler assigned values to the variables. What are these values called?
		code: here’s the solution: https://play.golang.org/p/jzHwSlles9
		video: 018*/

	fmt.Println(x)
	fmt.Println(y)
	fmt.Println(z)
}
