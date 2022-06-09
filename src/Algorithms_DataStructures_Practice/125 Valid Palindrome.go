package main

import (
	"fmt"
	"strings"
)

func main() {
	s := "agga"
	fmt.Println(isPalindrome(s))
}

func isPalindrome(s string) bool {
	s = strings.ToLower(s) // First we convert the input string to lowercase since the check is case insensitive as stated in the problem description
	// Then we are going to use two indices, one at the beginning of the string(left) starting at 0 and one at the end(right) of the string starting at len(s)-1
	// We are going to move both the indices at the same time: left will move forward and right will move backward (left, right = left+1, right-1)
	// If they cross (left < right), we know we're done since it means we have tested each pair of values left and right so far
	for left, right := 0, len(s)-1; left < right; left, right = left+1, right-1 {
		// Before testing if the left character is equal to the right character, we need to filter out the "invalid" characters using the isValid function
		// as long as the left character is invalid, we move the index forward. We also make sure that the indices don't cross or else it means it's not a palindrome
		for left < right && !isValid(s[left]) {
			left++
		}
		// same for the right pointer, as long as the right character is invalid, we move the index backward. We also make sure that the indices don't cross or else it means it's not a palindrome
		for left < right && !isValid(s[right]) {
			right--
		}
		// finally, we have two valid characters that we can compare. If they differ, it means it's not a palindrome, so return false immediately, if they match, just continue
		// We are moving both indices at the same time: left move forward and right move backward (left, right = left+1, right-1)
		if s[left] != s[right] {
			return false
		}
	}
	return true
}

// As stated in the problem description, "all non-alphanumeric characters" are considered invalid:
// So we return true only if a character is comprised into the range 'a' to 'z' or if it's a number,  comprised in the range '0' to '9'
// Note that we don't need to check for uppercase since we have converted the input string to lowercase earlier
func isValid(c byte) bool { return c >= 'a' && c <= 'z' || c >= '0' && c <= '9' }
