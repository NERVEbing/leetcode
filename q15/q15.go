package q15

import (
	"fmt"
	"sort"
)

func threeSum(nums []int) [][]int {
	var resp [][]int

	if len(nums) < 3 {
		return resp
	}

	sort.Ints(nums)

	for x := 0; x < len(nums); x++ {
		if x > 0 && nums[x] == nums[x-1] {
			continue
		}

		z := len(nums) - 1
		target := -nums[x]

		for y := x + 1; y < len(nums); y++ {
			if y > x+1 && nums[y] == nums[y-1] {
				continue
			}

			for y < z && nums[y]+nums[z] > target {
				z--
			}

			if y == z {
				break
			}

			if nums[y]+nums[z] == target {
				resp = append(resp, []int{nums[x], nums[y], nums[z]})
			}
		}
	}

	return resp
}

func Test() {
	nums := []int{-1, 0, 1, 2, -1, -4}
	result := threeSum(nums)
	fmt.Println(result)
}
