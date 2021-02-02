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

func isRectangleOverlap(rec1 []int, rec2 []int) bool {
	defer timeCost()()
	range:=func()bool{

	}

	if(range()||range()||range()){
		return true
	}
	return false
}
