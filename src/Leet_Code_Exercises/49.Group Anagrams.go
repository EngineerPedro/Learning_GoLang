package main

import "fmt"

func main() {
	strs := []string{"eat", "tea", "tan", "ate", "nat", "bat"}
	fmt.Println(groupAnagrams(strs))
}

func groupAnagrams(strs []string) [][]string {
	// First we create a map:
	//  - The key is [26]int, an array of integers which represents the number of occurrences of characters in
	// the string. Each index represents a letter (0 is 'a', 1 is 'b', ..., 25 is 'z'). For instance, for the
	//string "abcabe", we would have the array [2 2 1 0 1 .... 0 ]
	//  - The value is a slice of strings: we will associate all the strings with the same "signature" to the
	// same key in the map. For instance, "aba" and "baa" have the same signature [2 1 .... 0 ], so we would
	//associate them to the same key: { [2 1 ... 0 ]: []string{"aba", "baa" }
	groups := map[[26]int][]string{}
	// Then we want to populate that map
	for _, s := range strs { // for each string in the input
		key := [26]int{}
		for i := range s {
			key[s[i]-'a']++
		} // we compute the key (the signature) by counting
		// the numbers of occurrences of each characters in the string, the s[i]-'a' instruction maps any lowercase
		//character to the range [0 - 25]
		// once we have computed the key for the current string, we insert it in the map. Since the value is a
		//slice of strings, we just append the current string to the slice
		groups[key] = append(groups[key], s)
	}
	// after this we have all our anagram strings grouped together by key, but for the result, we don't want the
	//key, just a slice of "grouped anagrams"
	res := [][]string{} // so we create a slice of slice of string
	for _, v := range groups {
		res = append(res, v)
	} // and we range over our map to only keep the values
	// (we don't need the key anymore). We append the keys to the result
	return res //and we're done!
}
