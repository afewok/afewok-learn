package leetcode

import (
	"fmt"
	"testing"
)

/**
 * 1480. 一维数组的动态和
 */

func Test_leetcode_1480(t *testing.T) {
	fmt.Println(runningSum([]int{1, 2, 3, 4}))
	fmt.Println(runningSum([]int{1, 1, 1, 1, 1}))
	fmt.Println(runningSum([]int{3, 1, 2, 10, 1}))
}

func runningSum(nums []int) []int {
	defer timeCost()()
	length := len(nums)
	for i := 1; i < length; i++ {
		nums[i] = nums[i-1] + nums[i]
	}
	return nums
}
