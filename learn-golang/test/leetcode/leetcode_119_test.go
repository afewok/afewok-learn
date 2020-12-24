package leetcode

import (
	"fmt"
	"testing"
)

/**
 * 119. 杨辉三角 II
 */

func Test_leetcode_119(t *testing.T) {
	fmt.Println(getRow(3))
}

func getRow(rowIndex int) []int {
	if rowIndex == 0 {
		return []int{1}
	} else if rowIndex == 1 {
		return []int{1, 1}
	}
	history := getRow(rowIndex - 1)
	arr := make([]int, rowIndex+1)
	arr[0], arr[rowIndex] = 1, 1
	for i := 1; i < rowIndex; i++ {
		arr[i] = history[i-1] + history[i]
	}
	return arr
}
