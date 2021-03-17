package leetcode

import (
	"fmt"
	"testing"
)

/**
 * 695. 岛屿的最大面积
 *
 */
func Test_leetcode_695(t *testing.T) {
	fmt.Println(maxAreaOfIsland([][]int{
		{0, 0, 1, 0, 0, 0, 0, 1, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 1, 1, 1, 0, 0, 0},
		{0, 1, 1, 0, 1, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 1, 0, 0, 1, 1, 0, 0, 1, 0, 1, 0, 0},
		{0, 1, 0, 0, 1, 1, 0, 0, 1, 1, 1, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 1, 1, 1, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 1, 1, 0, 0, 0, 0},
	}))
}

func maxAreaOfIsland(grid [][]int) int {
	defer timeCost()()
	max := 0
	for y := 0; y < len(grid); y++ {
		for x := 0; x < len(grid[y]); x++ {
			if grid[y][x] == 1 {
				temp := getArea(grid, x, y)
				if max < temp {
					max = temp
				}
			}
		}
	}
	return max
}

func getArea(grid [][]int, x, y int) int {
	if y >= 0 && y < len(grid) && x >= 0 && x < len(grid[y]) && grid[y][x] == 1 {
		grid[y][x] = 0
		return 1 + getArea(grid, x-1, y) + getArea(grid, x+1, y) + getArea(grid, x, y-1) + getArea(grid, x, y+1)
	}
	return 0
}
