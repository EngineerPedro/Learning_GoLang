package main

import (
	"fmt"
)

func main() {
	//Exercise_01()
	//Exercise_02()
	//Exercise_03()
	//Exercise_04()
	//Exercise_05()
	//Exercise_06()
	//Exercise_07()
	//Exercise_08()
	//Exercise_09()
	Exercise_10()
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

func Exercise_02() {

	/*Hands-on exercise #2
		Print every rune code point of the uppercase alphabet three times. Your output should look like
	this:
		65
		U+0041 'A'
		U+0041 'A'
		U+0041 'A'
		66
		U+0042 'B'
		U+0042 'B'
		U+0042 'B'
		… through the rest of the alphabet characters
	solution: https://play.golang.org/p/1OjnCX1D5H
	video: 051*/

	for i := 65; i <= 90; i++ {
		fmt.Println(i)
		for j := 0; j < 3; j++ {
			fmt.Printf("\t%#U\n", i)
		}
	}

}

func Exercise_03() {
	/*Hands-on exercise #3
		Create a for loop using this syntax
		● for condition { }
		Have it print out the years you have been alive.
			solution: https://play.golang.org/p/tnyqBPJ-i5
	video: 052*/

	for x := 1993; x <= 2022; x++ {
		if x == 1993 {
			fmt.Println("I was born:", x)
		} else if x == 2022 {
			fmt.Println("And I am still alive this year :", x)
		} else {
			fmt.Println("I was alive :", x)
		}
	}
}

func Exercise_04() {
	/*Hands-on exercise #4
		Create a for loop using this syntax
		● for { }
		Have it print out the years you have been alive.
			solution: https://play.golang.org/p/9VpnB-I1Pz
	video: 053*/
	x := 1993
	for {
		if x == 1993 {
			fmt.Println("I was born:", x)
		} else if x == 2022 {
			fmt.Println("And I am still alive this year :", x)
			break
		} else {
			fmt.Println("I was alive :", x)
		}
		x++
	}

}

func Exercise_05() {

	/*Hands-on exercise #5
		Print out the remainder (modulus) which is found for each number between 10 and 100 when it
		is divided by 4.
	solution: https://play.golang.org/p/ohfJOW9euy
	video: 054*/

	x := 10
	for {
		if x == 10 {
			fmt.Println("Lets get the remainder of the numbers from 10-100 after dividing by 4")
			fmt.Println(x, " divided by 4 has a remainder of : ", x%4)
		} else if x == 100 {

			fmt.Println("And finally 100 divided by 4 is   :", x%4)
			break
		} else {
			if x%4 != 0 {
				fmt.Println(x, " divided by 4 has a remainder of : ", x%4)
			}
		}
		x++
	}
}

func Exercise_06() {
	/*	Hands-on exercise #6
		Create a program that shows the “if statement” in action.
			solution: https://play.golang.org/p/DpZ_FLfn5s
	video: 055*/

	//performed this in my preiviouse exercises b/c well
	//I know how to use an if statement already so just going to copy paste

	x := 10
	for {
		if x == 10 {
			fmt.Println("Lets get the remainder of the numbers from 10-100 after dividing by 4")
			fmt.Println(x, " divided by 4 has a remainder of : ", x%4)
		} else if x == 100 {

			fmt.Println("And finally 100 divided by 4 is   :", x%4)
			break
		} else {
			if x%4 != 0 {
				fmt.Println(x, " divided by 4 has a remainder of : ", x%4)
			}
		}
		x++
	}
}

func Exercise_07() {
	/*Hands-on exercise #7
		Building on the previous hands-on exercise, create a program that uses “else if” and “else”.
	solution: https://play.golang.org/p/IDnrJpE7vT
	video: 056*/
	// I have been combining for loop , if , and else if statements
	// This exercise I just copied over what I already did to move forward
	x := 1993
	for {
		if x == 1993 {
			fmt.Println("I was born:", x)
		} else if x == 2022 {
			fmt.Println("And I am still alive this year :", x)
			break
		} else {
			fmt.Println("I was alive :", x)
		}
		x++
	}
}

func Exercise_08() {
	/*	Hands-on exercise #8
		Create a program that uses a switch statement with no switch expression specified.
			solution: https://play.golang.org/p/YpPgkWGqKY
	video: 057*/
	x := 15
	y := 16
	switch {
	case x == y:
		fmt.Println("should not show ")
	case x > y:
		fmt.Println("should not show ")
	case x < y:
		fmt.Println("should show b/c y is greater then x and default " +
			"switch statement without expression goes to what is true ")

	}

}

func Exercise_09() {
	/*	Hands-on exercise #9
			Create a program that uses a switch statement with the switch expression specified as a
			variable of TYPE string with the IDENTIFIER “favSport”.
		solution: https://play.golang.org/p/h-8FnjpzWB
		video: 058*/
	favsport := "soccer"
	switch favsport {
	case "basketball":
		fmt.Println("shoot")
	case "soccer":
		fmt.Println("KICK")
	case "flying":
		fmt.Println("JUMP")

	}
}

func Exercise_10() {
	/*	Hands-on exercise #10
			Write down what these print:
			● fmt.Println(true && true)
			Todd McLeod - Learn To Code Go on Udemy - Part 1 - Page 38
			● fmt.Println(true && false)
			● fmt.Println(true || true)
			● fmt.Println(true || false)
			● fmt.Println(!true)
		solution:
		video: 059*/

	fmt.Println(true && true)
	fmt.Println(true && false)
	fmt.Println(true || true)
	fmt.Println(true || false)
	fmt.Println(!true)
}
