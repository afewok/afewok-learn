package leetcode

import (
	"fmt"
	"testing"
)

/**
 * 724. 寻找数组的中心索引
 *
 * 思路：
 */
func Test_leetcode_724(t *testing.T) {
	fmt.Println(pivotIndex([]int{-1, -1, -1, -1, -1, 0}))
	fmt.Println(pivotIndex([]int{1, 7, 3, 6, 5, 6}))
}

func pivotIndex(nums []int) int {
	defer timeCost()()
	length := len(nums)
	if length < 1 {
		return -1
	}
	numleft, numright := 0, 0
	for i := 0; i < length; i++ {
		numright += nums[i]
	}
	for i := 0; i < length; i++ {
		numright -= nums[i]
		if numleft == numright {
			return i
		}
		numleft += nums[i]
	}

	return -1
}
