package leetcode

import (
	"fmt"
	"testing"
)

/**
 * 74. 搜索二维矩阵
 */

func Test_leetcode_074(t *testing.T) {
	fmt.Println(searchMatrix([][]int{{1, 3}}, 3))
	fmt.Println(searchMatrix([][]int{{1, 3, 5, 7}, {10, 11, 16, 20}, {23, 30, 34, 60}}, 3))
	fmt.Println(searchMatrix([][]int{{1, 3, 5, 7}, {10, 11, 16, 20}, {23, 30, 34, 60}}, 13))
}

func searchMatrix(matrix [][]int, target int) bool {
	defer timeCost()()
	for _, items := range matrix {
		if items[len(items)-1] >= target {
			for _, item := range items {
				if item == target {
					return true
				}
			}
			return false
		}
	}
	return false
}
