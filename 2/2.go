package main

import (
	"fmt"
	"math"
	"sort"
)

func main() {
	fmt.Println(findMedianSortedArrays([]int{1, 3}, []int{2}))
}

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	slice := append(nums1, nums2...)
	sort.Ints(slice)
	midianaIndex := len(slice) / 2
	if len(slice)%2 == 0 {
		return (float64(slice[midianaIndex]) + float64(slice[midianaIndex-1])) / 2
	} else {
		return float64(slice[int64(math.Trunc(float64(midianaIndex)))])
	}
}
