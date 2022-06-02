package main

import "fmt"

func main() {
	var firstNumber int
	var secondNumber int

	fmt.Scanln(&firstNumber, &secondNumber)

	fmt.Println(secondNumber + firstNumber)
}
