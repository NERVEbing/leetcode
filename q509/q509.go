package q509

import "fmt"

func fib(n int) int {
	if n < 2 {
		return n
	}

	n1, n2, result := 0, 0, 1

	for i := 2; i <= n; i++ {
		n1 = n2
		n2 = result
		result = n1 + n2
	}

	return result
}

func Test() {
	n := 10

	for i := 0; i < n; i++ {
		result := fib(i)
		fmt.Printf("%d ", result)
	}
}
