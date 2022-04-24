package main

import "fmt"

func main() {
	//Exercise_01()
	Exercise_02()
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
