package q392

import (
	"fmt"
)

func isSubsequence(s string, t string) bool {
	if len(s) == 0 || s == t {
		return true
	}
	if len(s) > len(t) {
		return false
	}

	index := 0
	for i := 0; i < len(t); i++ {
		if s[index] == t[i] {
			index++
		}
		if len(s) == index {
			return true
		}
	}

	if len(s) == index {
		return true
	}

	return false
}

func Test() {
	s1 := "abc"
	s2 := "aebrc1"
	result := isSubsequence(s1, s2)
	fmt.Println(result)
}
