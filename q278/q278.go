package q278

import (
	"fmt"
)

func firstBadVersion(n int) int {
	min := 1
	max := n

	for min < max {
		mid := min + (max-min)/2

		if isBadVersion(mid) {
			max = mid
		} else {
			min = mid + 1
		}
	}

	return max
}

func isBadVersion(version int) bool {
	return version > 20
}

func Test() {
	n := 300
	result := firstBadVersion(n)
	fmt.Println(result)
}
