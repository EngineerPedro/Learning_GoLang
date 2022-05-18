package main

import (
	"encoding/json"
	"fmt"
	"os"
	"sort"
)

func main() {
	//Exercise_01()
	//Exercise_02()
	//Exercise_03()
	//Exercise_04()
	Exercise_05()

}

func Exercise_01() {
	/*
		Hands-on exercise #1
		Starting with this code, marshal the []user to JSON. There is a little bit of a curve ball in this
		hands-on exercise - remember to ask yourself what you need to do to EXPORT a value from a
		package.
		solution: https://play.golang.org/p/8BK6PXj3aG
		video: 125
	*/

	u1 := user{
		First: "James",
		Age:   32,
	}

	u2 := user{
		First: "Moneypenny",
		Age:   27,
	}

	u3 := user{
		First: "M",
		Age:   54,
	}

	users := []user{u1, u2, u3}

	fmt.Println(users)

	// your code goes here
	fmt.Println(" ")
	bs, err := json.Marshal(users)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(bs))
}

func Exercise_02() {
	/*
		Hands-on exercise #2
			Starting with this code, unmarshal the JSON into a Go data structure. Hint:
			https://mholt.github.io/json-to-go/
			code:
			● code setup (just fyi, not needed for exercise):
			○ https://play.golang.org/p/nWPP5Z2Q4e
			● solution:
			○ https://play.golang.org/p/r8oSG8DIPR
			video: 126

	*/

	s := `[{"First":"James","Last":"Bond","Age":32,"Sayings":["Shaken, not stirred",
			"Youth is no guarantee of innovation",
			"In his majesty's royal service"]},{"First":"Miss","Last":"Moneypenny",
			"Age":27,"Sayings":["James, it is soo good to see you","Would you like me to take care of that for you, James?",
			"I would really prefer to be a secret agent myself."]},{"First":"M","Last":"Hmmmm",
			"Age":54,"Sayings":["Oh, James. You didn't.","Dear God, what has James done now?",
			"Can someone please tell me where James Bond is?"]}]`
	fmt.Println(s)

	var people []person

	err := json.Unmarshal([]byte(s), &people)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(people)

	for i, person := range people {
		fmt.Println("Person #", i)
		fmt.Println("\t", person.First, person.Last, person.Age)
		for _, saying := range person.Sayings {
			fmt.Println("\t\t", saying)
		}
	}
}

func Exercise_03() {
	/*
		Hands-on exercise #3
		Starting with this code, encode to JSON the []user sending the results to Stdout. Hint: you will
		need to use json.NewEncoder(os.Stdout).encode(v interface{})
		solution: https://play.golang.org/p/ql90D1OQqd
		video: 127

	*/
	u1 := user2{
		First: "Pedro",
		Last:  "Guzman",
		Age:   29,
		Sayings: []string{
			"Shaken, not stirred",
			"Youth is no guarantee of innovation",
			"In his majesty's royal service",
		},
	}

	u2 := user2{
		First: "Malaysia",
		Last:  "Guzman",
		Age:   27,
		Sayings: []string{
			"James, it is soo good to see you",
			"Would you like me to take care of that for you, James?",
			"I would really prefer to be a secret agent myself.",
		},
	}

	u3 := user2{
		First: "Leslie",
		Last:  "Guzman",
		Age:   30,
		Sayings: []string{
			"Oh, James. You didn't.",
			"Dear God, what has James done now?",
			"Can someone please tell me where James Bond is?",
		},
	}

	users := []user2{u1, u2, u3}

	fmt.Println(users)

	err := json.NewEncoder(os.Stdout).Encode(users)
	if err != nil {
		fmt.Println("We did something wrong and here's the error:", err)
	}

}

func Exercise_04() {
	xi := []int{5, 8, 2, 43, 17, 987, 14, 12, 21, 1, 4, 2, 3, 93, 13}
	xs := []string{"random", "rainbow", "delights", "in", "torpedo", "summers", "under", "gallantry", "fragmented", "moons", "across", "magenta"}

	fmt.Println(xi)
	// sort xi
	sort.Ints(xi)
	fmt.Println(xi)

	fmt.Println("###################################")
	fmt.Println("###################################")

	fmt.Println(xs)
	// sort xs
	sort.Strings(xs)
	fmt.Println(xs)
}

func Exercise_05() {
	/*
		Hands-on exercise #5
		Starting with this code, sort the []user by
		● age
		● last
		Also sort each []string “Sayings” for each user
		● print everything in a way that is pleasant
		solution: https://play.golang.org/p/8RKkdtLl6w
		video: 129
	*/

	u1 := user3{
		First: "James",
		Last:  "Bond",
		Age:   32,
		Sayings: []string{
			"Shaken, not stirred",
			"Youth is no guarantee of innovation",
			"In his majesty's royal service",
		},
	}

	u2 := user3{
		First: "Miss",
		Last:  "Moneypenny",
		Age:   27,
		Sayings: []string{
			"James, it is soo good to see you",
			"Would you like me to take care of that for you, James?",
			"I would really prefer to be a secret agent myself.",
		},
	}

	u3 := user3{
		First: "M",
		Last:  "Hmmmm",
		Age:   54,
		Sayings: []string{
			"Oh, James. You didn't.",
			"Dear God, what has James done now?",
			"Can someone please tell me where James Bond is?",
		},
	}

	users := []user3{u1, u2, u3}

	fmt.Println(users)

	// your code goes here

	for _, u := range users {
		fmt.Println(u.First, u.Last, u.Age)
		for _, v := range u.Sayings {
			fmt.Println("\t", v)
		}
	}

	sort.Sort(ByAge(users))

	for _, u := range users {
		fmt.Println(u.First, u.Last, u.Age)
		for _, v := range u.Sayings {
			fmt.Println("\t", v)
		}
	}
}

//Work Area

type user struct {
	First string
	Age   int
}

type person struct {
	First   string   `json:"First"`
	Last    string   `json:"Last"`
	Age     int      `json:"Age"`
	Sayings []string `json:"Sayings"`
}

type user2 struct {
	First   string
	Last    string
	Age     int
	Sayings []string
}

type user3 struct {
	First   string
	Last    string
	Age     int
	Sayings []string
}

type ByAge []user3

func (a ByAge) Len() int           { return len(a) }
func (a ByAge) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByAge) Less(i, j int) bool { return a[i].Age < a[j].Age }

type ByLast []user3

func (l ByLast) Len() int           { return len(l) }
func (l ByLast) Swap(i, j int)      { l[i], l[j] = l[j], l[i] }
func (l ByLast) Less(i, j int) bool { return l[i].Last < l[j].Last }
