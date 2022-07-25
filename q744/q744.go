package q744

import (
	"fmt"
)

func nextGreatestLetter(letters []byte, target byte) byte {
	min := 0
	max := len(letters) - 1

	if target >= letters[max] {
		return letters[min]
	}

	for min < max {
		mid := min + (max-min)/2
		if letters[mid] > target {
			max = mid
		} else {
			min = mid + 1
		}
	}

	return letters[min]
}

func Test() {
	letters := []byte{'c', 'f', 'j'}
	target := byte('j')
	result := nextGreatestLetter(letters, target)
	fmt.Println(string(result))
}
