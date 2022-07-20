package q189

func rotate(nums []int, k int) {
	newNums := make([]int, len(nums))
	for n, m := range nums {
		newNums[(n+k)%len(nums)] = m
	}

	copy(nums, newNums)
}

func Test() {
	nums := []int{-1, -100, 3, 99}
	k := 2
	rotate(nums, k)
}
