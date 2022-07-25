package q69

import (
	"fmt"
)

func mySqrt(x int) int {
	min := 0
	max := x
	result := -1

	for min <= max {
		mid := min + (max-min)/2
		if mid*mid <= x {
			result = mid
			min = mid + 1
		} else {
			max = mid - 1
		}
	}

	return result
}

func Test() {
	x := 101
	result := mySqrt(x)
	fmt.Println(result)
}
