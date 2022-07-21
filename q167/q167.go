package q167

import (
	"fmt"
)

func twoSum(numbers []int, target int) []int {
	left := 0
	right := len(numbers) - 1

	for left < right {
		if numbers[left]+numbers[right] == target {
			return []int{left + 1, right + 1}
		}
		if numbers[left]+numbers[right] < target {
			left++
		} else {
			right--
		}
	}

	return nil
}

func Test() {
	nums := []int{2, 7, 11, 15}
	target := 18
	result := twoSum(nums, target)
	fmt.Println(result)
}
