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
	Exercise_09and10()
	//Practicing()
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

func Exercise_04() {
	/*Hands-on exercise #4
		Follow these steps:
		● start with this slice
		○ x := []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51}
		● append to that slice this value
		○ 52
		● print out the slice
		● in ONE STATEMENT append to that slice these values
		○ 53
		○ 54
		○ 55
		● print out the slice
		● append to the slice this slice
		○ z := []int{56, 57, 58, 59, 60}
		● print out the slice
	solution: https://play.golang.org/p/QUYhtSBaDS
	video: 074*/

	x := []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51}
	x = append(x, 52)
	fmt.Println(x)
	y := []int{53, 54, 55}
	x = append(x, y...)
	fmt.Println(x)
	z := []int{56, 57, 58, 59, 60}
	x = append(x, z...)
	fmt.Println(x)
}

func Exercise_05() {
	/*	Hands-on exercise #5
		To DELETE from a slice, we use APPEND along with SLICING. For this hands-on exercise,
			follow these steps:
		● start with this slice
		○ x := []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51}
		● use APPEND & SLICING to get these values here which you should ASSIGN to a
		variable “y” and then print:
		○ [42, 43, 44, 48, 49, 50, 51]
		solution: https://play.golang.org/p/u8zpHLfgSE
		video: 075*/

	x := []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51}
	y := append(x[:3], x[6:]...)
	fmt.Println(y)

}

func Exercise_06() {
	/*Hands-on exercise #6
	Create a slice to store the names of all of the states in the United States of America. Use make
	and append to do this. Goal: do not have the array that underlies the slice created more than
	once. What is the length of your slice? What is the capacity? Print out all of the values, along
	with their index position, without using the range clause. Here is a list of the 50 states:*/

	x := make([]string, 50, 100)

	y := []string{`Alabama`, `Alaska`, `Arizona`, `Arkansas`, `California`, `Colorado`, `Connecticut`,
		`Delaware`, `Florida`, `Georgia`, `Hawaii`, `Idaho`, `Illinois`, `Indiana`, `Iowa`, `Kansas`, `Kentucky`,
		`Louisiana`, `Maine`, `Maryland`, `Massachusetts`, `Michigan`, `Minnesota`, `Mississippi`, `Missouri`, `Montana`,
		`Nebraska`, `Nevada`, `New Hampshire`, `New Jersey`, `New Mexico`, `New York`, `North Carolina`, `North Dakota`,
		`Ohio`, `Oklahoma`, `Oregon`, `Pennsylvania`, `Rhode Island`, `South Carolina`, `South Dakota`, `Tennessee`, `Texas`,
		`Utah`, `Vermont`, `Virginia`, `Washington`, `West Virginia`, `Wisconsin`, `Wyoming`}

	for i := 0; i <= 49; i++ {
		x[i] = y[i]
		fmt.Println(i, x[i])
	}

	fmt.Println(x)
	fmt.Println(len(x))
	fmt.Println(cap(x))
}

func Exercise_07() {
	/*Hands-on exercise #7
		Create a slice of a slice of string ([][]string). Store the following data in the multi-dimensional
	slice:
		"James", "Bond", "Shaken, not stirred"
		"Miss", "Moneypenny", "Helloooooo, James."
		Range over the records, then range over the data in each record.
			solution: https://play.golang.org/p/FMM4c2PhZg
	video: 077*/

	xs1 := []string{"James", "Bond", "Shaken, not stirred"}
	xs2 := []string{"Miss", "Moneypenny", "Helloooooo, James."}
	fmt.Println(xs1)
	fmt.Println(xs2)

	xxs := [][]string{xs1, xs2}
	fmt.Println(xxs)

	for i, xs := range xxs {
		fmt.Println("record: ", i)
		for j, val := range xs {
			fmt.Printf("\t index position: %v \t value: \t %v \n", j, val)
		}
	}

}

func Exercise_08() {
	/*	Hands-on exercise #8
			Create a map with a key of TYPE string which is a person’s “last_first” name, and a value of
			TYPE []string which stores their favorite things. Store three records in your map. Print out all of
			the values, along with their index position in the slice.
			`bond_james`, `Shaken, not stirred`, `Martinis`, `Women`
			`moneypenny_miss`, `James Bond`, `Literature`, `Computer Science`
			`no_dr`, `Being evil`, `Ice cream`, `Sunsets`
		solution: https://play.golang.org/p/nTzSlRa9_A
		video: 078*/
	m := map[string][]string{
		"Guzman_Pedro":  []string{"Coding", "Arguing", "Reading"},
		"Guzman_Leslie": []string{"Cooking", "Running", "Yardword"},
		"Dog_Sophie":    []string{"Sleeping", "Eating", "Pooping"},
	}
	for i, v := range m {

		println(m)
	}
}

func Exercise_08() {
	m := map[string][]string{
		`bond_james`:      []string{`Shaken, not stirred`, `Martinis`, `Women`},
		"moneypenny_miss": []string{"James Bond", "Literature", "Computer Science"},
		"Agent_Pedro":     []string{`Being evil`, `Ice cream`, `Sunsets`},
	}
	fmt.Println(m)

	for k, v := range m {
		fmt.Println("This is the record for:", k)
		for i, v2 := range v {
			fmt.Println("\t", i, v2)
		}
	}
}

func Exercise_09and10() {
	m := map[string][]string{
		`bond_james`:      []string{`Shaken, not stirred`, `Martinis`, `Women`},
		"moneypenny_miss": []string{"James Bond", "Literature", "Computer Science"},
		"Agent_Pedro":     []string{`Being evil`, `Ice cream`, `Sunsets`},
		"Agent_Leslie":    []string{"Pedro", "Running", "Home-Maker"},
	}
	fmt.Println(m)

	delete(m, "moneypenny_miss")

	for k, v := range m {
		fmt.Println("This is the record for:", k)
		for i, v2 := range v {
			fmt.Println("\t", i, v2)
		}
	}

	delete(m, "moneypenny_miss")
}
func Practicing() {
	//Declare 4 slices and print each of them to embed the process
	_Agent_Pedro := []string{"coder", "Computer-Scientists", "pyro", "carnivore"}
	fmt.Println(_Agent_Pedro)
	fmt.Printf("%T\n", _Agent_Pedro)
	_Anget_Leslie := []string{"home-maker", "Cooks", "perfectionist", "carnivore"}
	fmt.Println(_Anget_Leslie)
	fmt.Printf("%T\n", _Anget_Leslie)
	_Agent_Andrew := []string{"gun-smith", "Hunter", "debater", "carnivore"}
	fmt.Println(_Agent_Andrew)
	fmt.Printf("%T\n", _Agent_Andrew)
	_Agent_Rachel := []string{"stay-at-home-mom", "Writer", "Author", "carnivore"}
	fmt.Println(_Agent_Rachel)
	fmt.Printf("%T\n", _Agent_Rachel)

	//Chopping the slices up
	_Agent_Pedro = _Agent_Pedro[1:]
	fmt.Println(_Agent_Pedro)
	fmt.Printf("%T\n", _Agent_Pedro)

	_Agent_Pedro = _Agent_Pedro[:2]
	fmt.Println(_Agent_Pedro)
	fmt.Printf("%T\n", _Agent_Pedro)

}
