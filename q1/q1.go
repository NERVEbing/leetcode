package q1

import "fmt"

func twoSum(nums []int, target int) []int {
	for k, v := range nums {
		other := target - v
		for i := k + 1; i < len(nums); i++ {
			if other == nums[i] {
				return []int{k, i}
			}
		}
	}
	return nil
}

func Test() {
	result := twoSum([]int{3, 2, 4}, 6)
	fmt.Println(result)
}
