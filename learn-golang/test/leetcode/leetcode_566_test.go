package leetcode

import (
	"fmt"
	"testing"
)

/**
 * 566. 重塑矩阵
 */

func Test_leetcode_566(t *testing.T) {
	fmt.Println(matrixReshape([][]int{{1, 2}, {3, 4}}, 1, 4))
	fmt.Println(matrixReshape([][]int{{1, 2}, {3, 4}}, 2, 4))
}

func matrixReshape(nums [][]int, r int, c int) [][]int {
	defer timeCost()()
	height, weight := len(nums), len(nums[0])
	if height*weight != r*c {
		return nums
	}
	h1, w1, h2, w2, result := 0, 0, 0, 0, make([][]int, r)
	for h1 < height {
		if w2 == 0 {
			result[h2] = make([]int, c)
		}
		result[h2][w2] = nums[h1][w1]
		w1++
		w2++
		if w1 == weight {
			w1 = 0
			h1++
		}
		if w2 == c {
			w2 = 0
			h2++
		}
	}
	return result
}
