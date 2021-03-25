package leetcode

import (
	"fmt"
	"testing"
)

/**
 * 977. 有序数组的平方
 */
func Test_leetcode_977(t *testing.T) {
	fmt.Println(sortedSquares([]int{-4, -1, 0, 3, 10}))
	fmt.Println(sortedSquares([]int{-7, -3, 2, 3, 11}))
}

func sortedSquares(nums []int) []int {
	defer timeCost()()
	length := len(nums)
	for i := 0; i < length; i++ {
		nums[i] = nums[i] * nums[i]
	}
	for i := 0; i < length; i++ {
		for j := i + 1; j < length; j++ {
			if nums[i] > nums[j] {
				nums[i], nums[j] = nums[j], nums[i]
			}
		}
	}
	return nums
}
