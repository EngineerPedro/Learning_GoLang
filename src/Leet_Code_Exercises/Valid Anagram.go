package main

import "fmt"

func main() {
	s := "anagram"
	t := "nagaram"
	fmt.Println(isAnagram(s, t))

}

func isAnagram(s string, t string) bool {
	//first check if te length of both of the strings is the same
	if len(s) != len(t) {
		return false
	}

	//create a hash map whos keys will be the letter
	//and number of times the letter occurs is the
	//value
	hm := make(map[rune]int)

	//adds the key and value (number of times it occurs to that key)
	for _, index2 := range s {
		hm[index2]++
	}

	//redo the same for the second string BUT in reverse
	// this will reduce every value by each key to zero
	//IF these strings are in fact anagrams of one another
	for _, index1 := range t {
		hm[index1]--
	}

	//Go through all values of the hashmap
	//makes sure that they are ZERO and return true
	for k, _ := range hm {
		if hm[k] != 0 {
			return false
		}
	}
	//If each key value is not at 0
	return true
}
