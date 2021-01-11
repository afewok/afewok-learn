package leetcode

import (
	"fmt"
	"testing"
)

/**
 * 1203. 项目管理
 */

func Test_leetcode_1203(t *testing.T) {
	fmt.Println(sortItems(8, 2, []int{-1, -1, 1, 0, 0, 1, 0, -1}, [][]int{{}, {6}, {5}, {6}, {3, 6}, {}, {}, {}}))
}

func sortItems(n int, m int, group []int, beforeItems [][]int) []int {
	defer timeCost()()

}
