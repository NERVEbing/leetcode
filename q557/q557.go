package q557

import (
	"fmt"
)

func reverseWords(s string) string {
	var b []byte
	l := len(s)

	for i := 0; i < l; {
		head := i
		for i < l && s[i] != ' ' {
			i++
		}
		for p := head; p < i; p++ {
			b = append(b, s[head+i-1-p])
		}
		for i < l && s[i] == ' ' {
			i++
			b = append(b, ' ')
		}
	}

	return string(b)
}

func Test() {
	s := "Let's take LeetCode contest"
	result := reverseWords(s)
	fmt.Println(result)
}
