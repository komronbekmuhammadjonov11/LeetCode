package main

import (
	"fmt"
)

func main() {
	fmt.Println(longestPalindrome("qwfgwdabaqw"))
}

func longestPalindrome(s string) string {
	var longestStr string
	for i := 0; i < len(s); i++ {
		for j := 0; j <= i; j++ {

		}
	}
	return longestStr
}

func isPalindrome(s string) bool {
	if len(s) == 0 {
		return false
	}
	lastIndex := len(s) - 1
	for index := range s {
		if s[index] != s[lastIndex-index] {
			return false
		}
	}
	return true
}
