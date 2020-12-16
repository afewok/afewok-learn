package leetcode

import (
	"fmt"
	"testing"
)

/**
 * 48. 旋转图像
 *
 * 思路：
 */
func Test_leetcode_048(t *testing.T) {
	data1 := [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}
	data2 := [][]int{{5, 1, 9, 11}, {2, 4, 8, 10}, {13, 3, 6, 7}, {15, 14, 12, 16}}
	rotate(data1)
	rotate(data2)
	fmt.Println(data1)
	fmt.Println(data2)
}

func rotate(matrix [][]int) {
	defer timeCost()()
	length := len(matrix)
	for i := 0; i < length; i++ {

	}
}
