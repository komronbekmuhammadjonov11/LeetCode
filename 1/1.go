package main

import "fmt"

func main() {
	arr := []int{1, 4, 5}
	fmt.Println(twoSum(arr, 9))
}

func twoSum(nums []int, target int) []int {

	for i := 0; i < len(nums); i++ {
		for j := 0; j < len(nums); j++ {
			if nums[i]+nums[j] == target && i != j {
				var result = []int{i, j}
				return result
			}
		}
	}
	return []int{}
}
