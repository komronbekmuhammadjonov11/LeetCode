package main

import (
	"fmt"
)

func main() {
	fmt.Println(longestPalindrome("aziza"))
}

func longestPalindrome(s string) string {
	if len(s) == 1 {
		return s
	}
	var longestStr string
	for i := 0; i < len(s); i++ {
		for j := 0; j < len(s); j++ {
			if i < j && isPalindrome(s[i:j+1]) {
				if len(s[i:j+1]) > len(longestStr) {
					longestStr = s[i : j+1]
				}
			}
		}
	}
	if len(longestStr) == 0 {
		return s[0:1]
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
