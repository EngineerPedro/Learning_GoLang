package main

import "fmt"

func main() {
	//Exercise_01()
	//Exercise_02()
	//Exercise_03()
	Exercise_04()
}

func Exercise_01() {
	/*Hands-on exercise #1
	Create your own type “person” which will have an underlying type of “struct” so that it can store the following data:
		● first name
		● last name
		● favorite ice cream flavors
		Create two VALUES of TYPE person. Print out the values, ranging over the elements in the slice which stores the favorite flavors.
		solution:
		● https://play.golang.org/p/ouRHmH_POg*/
	type person struct {
		first      string
		last       string
		favFlavors []string
	}
	Pedro := person{
		first:      "Pedro",
		last:       "Guzman",
		favFlavors: []string{"straberry", "vanilla", "cakebatter"},
	}

	fmt.Println(Pedro.first)
	fmt.Println(Pedro.last)
	for i, v := range Pedro.favFlavors {
		fmt.Println(i, v)
	}
}

func Exercise_02() {
	/*	Hands-on exercise #2
		Take the code from the previous exercise, then store the values of type person in a map with the key of last name.
			Access each value in the map. Print out the values, ranging over the slice. solution: https://play.golang.org/p/RvvLyAMvGo*/
	type person struct {
		first      string
		last       string
		favFlavors []string
	}
	p1 := person{
		first: "James",
		last:  "Bond",
		favFlavors: []string{
			"chocolate",
			"martini",
			"rum and coke",
		},
	}

	p2 := person{
		first: "Miss",
		last:  "Moneypenny",
		favFlavors: []string{
			"strawberry",
			"vanilla",
			"capuccino",
		},
	}

	m := map[string]person{
		p1.last: p1,
		p2.last: p2,
	}

	for _, v := range m {
		fmt.Println(v.first)
		fmt.Println(v.last)
		for i, val := range v.favFlavors {
			fmt.Println(i, val)
		}
		fmt.Println("-------")
	}
}

func Exercise_03() {
	/*	Hands-on exercise #3
		● Create a new type: vehicle.
		○ The underlying type is a struct. ○ The fields:
			■ doors
			■ color
			● Create two new types: truck & sedan.
			○ The underlying type of each of these new types is a struct.
			○ Embed the “vehicle” type in both truck & sedan.
			○ Give truck the field “fourWheel” which will be set to bool.
			○ Give sedan the field “luxury” which will be set to bool. solution
			● Using the vehicle, truck, and sedan structs:
			○ using a composite literal, create a value of type truck and assign values to the
			fields;
			○ using a composite literal, create a value of type sedan and assign values to the
			fields.
			● Print out each of these values.
			● Print out a single field from each of these values.
			solution: https://play.golang.org/p/PrTtTv_vVO
	*/

	type vehicle struct {
		doors int
		color string
	}

	type truck struct {
		vehicle
		fourwheel bool
	}

	type sedan struct {
		vehicle
		luxury bool
	}

	t := truck{
		vehicle: vehicle{
			doors: 2,
			color: "white",
		},
		//fourWheel: true,
	}

	s := sedan{
		vehicle: vehicle{
			doors: 4,
			color: "black",
		},
		luxury: false,
	}
	fmt.Println(t)
	fmt.Println(s)
	fmt.Println(t.doors)
	fmt.Println(s.doors)
}

func Exercise_04() {
	/*	Hands-on exercise #4 Create and use an anonymous struct.
		solution: https://play.golang.org/p/otBHFs-lPp video: 089
	*/
	s := struct {
		first     string
		friends   map[string]int
		favDrinks []string
	}{
		first: "James",
		friends: map[string]int{
			"Moneypenny": 555,
			"Q":          777,
			"M":          888,
		},
		favDrinks: []string{
			"Martini",
			"Water",
		},
	}
	fmt.Println(s.first)
	fmt.Println(s.friends)
	fmt.Println(s.favDrinks)

	for k, v := range s.friends {
		fmt.Println(k, v)
	}

	for i, val := range s.favDrinks {
		fmt.Println(i, val)
	}

}
