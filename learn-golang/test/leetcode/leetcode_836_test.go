package leetcode

import (
	"fmt"
	"testing"
)

/**
 * 836. 矩形重叠
 */

func Test_leetcode_836(t *testing.T) {
	fmt.Println(isRectangleOverlap([]int{0, 0, 2, 2}, []int{1, 1, 3, 3}))
	fmt.Println(isRectangleOverlap([]int{0, 0, 1, 1}, []int{1, 0, 2, 1}))
	fmt.Println(isRectangleOverlap([]int{0, 0, 1, 1}, []int{2, 2, 3, 3}))
}

// x1<x2
// y1<y2
func isRectangleOverlap(rec1 []int, rec2 []int) bool {
	defer timeCost()()
	if rec1[0] >= rec2[2] || rec2[0] >= rec1[2] || rec1[1] >= rec2[3] || rec2[1] >= rec1[3] {
		return false
	} else if rec1[0] == rec1[2] || rec1[1] == rec1[3] || rec2[0] == rec2[2] || rec2[1] == rec2[3] {
		return false
	}
	return true
}
