package q20

import "fmt"

func isValid(s string) bool {
	if s == "" {
		return true
	}

	m := map[uint8]uint8{
		')': '(',
		'}': '{',
		']': '[',
	}

	var temp []uint8

	for i := 0; i < len(s); i++ {
		if s[i] == '(' || s[i] == '{' || s[i] == '[' {
			temp = append(temp, s[i])
		} else {
			if len(temp) == 0 {
				return false
			}
			if temp[len(temp)-1] != m[s[i]] {
				return false
			}
			temp = temp[:len(temp)-1]
		}
	}

	return len(temp) == 0
}

func Test() {
	result := isValid("([])")
	fmt.Println(result)
}
