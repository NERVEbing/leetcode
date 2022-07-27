package q200

import "fmt"

func numIslands(grid [][]byte) int {
	var f func(int, int)
	count := 0
	flag := byte('1')

	f = func(x int, y int) {
		xLen := len(grid)
		yLen := len(grid[x])

		grid[x][y] = '0'

		if x-1 >= 0 {
			if grid[x-1][y] == flag {
				f(x-1, y)
			}
		}

		if x+1 < xLen {
			if grid[x+1][y] == flag {
				f(x+1, y)
			}
		}

		if y-1 >= 0 {
			if grid[x][y-1] == flag {
				f(x, y-1)
			}
		}

		if y+1 < yLen {
			if grid[x][y+1] == flag {
				f(x, y+1)
			}
		}
	}

	for x := 0; x < len(grid); x++ {
		for y := 0; y < len(grid[x]); y++ {
			if grid[x][y] == flag {
				f(x, y)
				count++
			}
		}
	}

	return count
}

func Test() {
	grid := [][]byte{
		{'1', '1', '1', '1', '0'},
		{'1', '1', '0', '1', '0'},
		{'1', '1', '0', '0', '0'},
		{'0', '0', '0', '0', '0'},
	}
	result := numIslands(grid)
	fmt.Println(result)
}
