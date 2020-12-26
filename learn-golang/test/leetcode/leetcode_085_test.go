package leetcode

import (
	"fmt"
	"testing"
)

/**
 * 85. 最大矩形
 */

func Test_leetcode_085(t *testing.T) {
	fmt.Println(maximalRectangle([][]byte{{'1', '0', '1', '0', '0'}, {'1', '0', '1', '1', '1'}, {'1', '1', '1', '1', '1'}, {'1', '0', '0', '1', '0'}}))
	fmt.Println(maximalRectangle([][]byte{}))
	fmt.Println(maximalRectangle([][]byte{{'0'}}))
	fmt.Println(maximalRectangle([][]byte{{'1'}}))
	fmt.Println(maximalRectangle([][]byte{{'0', '0'}}))
}

func maximalRectangle(matrix [][]byte) int {
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return 0
	}
	height, result := make([]int, len(matrix[0])), 0
	for i := 0; i < len(matrix); i++ {
		for j := range height {
			height[j] = (height[j] + 1) * int(matrix[i][j]-'0')

			if height[j] > 0 {
				minHeight := height[j]
				for k := j; k >= 0 && height[k] > 0; k-- {
					if minHeight > height[k] {
						minHeight = height[k]
					}
					s := minHeight * (j - k + 1)

					if result < s {
						result = s
					}
				}
			}
		}
	}
	return result
}
