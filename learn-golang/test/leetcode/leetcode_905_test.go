package leetcode

import (
	"fmt"
	"testing"
)

/**
 * 905. 按奇偶排序数组
 */
func Test_leetcode_905(t *testing.T) {
	fmt.Println(sortArrayByParity([]int{3, 1, 2, 4}))
}

func sortArrayByParity(A []int) []int {
	defer timeCost()()
	left, right := make([]int, 0), make([]int, 0)
	for _, v := range A {
		if v%2 == 0 {
			left = append(left, v)
		} else {
			right = append(right, v)
		}
	}
	return append(left, right...)
}
