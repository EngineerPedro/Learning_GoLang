package main

import "fmt"

func main() {
	strs := []string{"eat", "tea", "tan", "ate", "nat", "bat"}
	fmt.Println(groupAnagrams(strs))
}

func groupAnagrams(strs []string) [][]string {
	dict := make(map[[26]int][]string)
	for _, v := range strs {
		ana := [26]int{}
		for _, c := range v {
			ana[c-'a']++
		}

		if _, ok := dict[ana]; !ok {
			dict[ana] = make([]string, 0)
		}
		dict[ana] = append(dict[ana], v)
	}

	res := make([][]string, 0)
	for _, v := range dict {
		res = append(res, v)
	}
	return res
}
