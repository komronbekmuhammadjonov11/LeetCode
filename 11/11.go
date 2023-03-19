package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(maxArea([]int{2, 4}))
}

func maxArea(height []int) int {
	var max int = 0

	for i := 0; i < len(height); i++ {
		for j := 0; j < len(height); j++ {
			if i >= j {
				continue
			}
			var minNum int = int(math.Min(float64(height[i]), float64(height[j])))
			if (j-i)*minNum > max {
				max = (j - i) * minNum
			}
		}
	}
	return max

}
