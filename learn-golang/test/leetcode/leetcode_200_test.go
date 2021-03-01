package leetcode

import (
	"fmt"
	"testing"
)

/**
 * 200. 岛屿数量
 */

func Test_leetcode_200(t *testing.T) {
	fmt.Println(numIslands([][]byte{
		{'1', '1', '1', '1', '0'},
		{'1', '1', '0', '1', '0'},
		{'1', '1', '0', '0', '0'},
		{'0', '0', '0', '0', '0'},
	}))

	fmt.Println(numIslands([][]byte{
		{'1', '1', '0', '0', '0'},
		{'1', '1', '0', '0', '0'},
		{'0', '0', '1', '0', '0'},
		{'0', '0', '0', '1', '1'},
	}))
}

func numIslands(grid [][]byte) int {
	defer timeCost()()
	result, lenH, lenW := 0, len(grid), len(grid[0])
	for h := 0; h < lenH; h++ {
		for w := 0; w < lenW; w++ {
			if grid[h][w] == '1' {
				dfs200(grid, h, w, lenH, lenW)
				result++
			}
		}
	}
	return result
}

func dfs200(grid [][]byte, h, w int, lenH, lenW int) {
	if h >= 0 && w >= 0 && h < lenH && w < lenW && grid[h][w] == '1' {
		grid[h][w] = '0'
		dfs200(grid, h-1, w, lenH, lenW)
		dfs200(grid, h+1, w, lenH, lenW)
		dfs200(grid, h, w-1, lenH, lenW)
		dfs200(grid, h, w+1, lenH, lenW)
	}
}
