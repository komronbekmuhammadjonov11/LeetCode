package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println(isPalindrome(-202))
}

func isPalindrome(num int) bool {
	s := strconv.Itoa(num)
	fmt.Println(s)
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
