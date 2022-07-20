package q977

import (
	"fmt"
)

func sortedSquares(nums []int) []int {
	n := make([]int, len(nums))
	low := 0
	high := len(nums) - 1

	for pos := len(nums) - 1; pos >= 0; pos-- {
		x := nums[low] * nums[low]
		y := nums[high] * nums[high]
		if x > y {
			n[pos] = x
			low++
		} else {
			n[pos] = y
			high--
		}
	}

	return n
}

func Test() {
	nums := []int{-7, -3, 2, 3, 11}
	result := sortedSquares(nums)
	fmt.Println(result)
}
