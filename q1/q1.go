package q1

import "fmt"

func twoSum(nums []int, target int) []int {
	m := make(map[int]int)
	for i := 0; i < len(nums); i++ {
		other := target - nums[i]
		if f, ok := m[other]; ok {
			return []int{f, i}
		}
		m[nums[i]] = i
	}

	return nil
}

func Test() {
	result := twoSum([]int{3, 2, 4}, 6)
	fmt.Println(result)
}
