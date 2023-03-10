package main

import (
	"fmt"
)

func main() {

	fmt.Println(reverse(-123))
}

func reverse(num int) int32 {
	res := 0
	for num != 0 {
		remainder := num % 10
		res = (res * 10) + remainder
		num /= 10
	}
	return int32(res)
}
