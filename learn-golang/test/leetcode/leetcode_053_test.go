package leetcode

import (
	"fmt"
	"testing"
)

/**
 * 53. 最大子序和
 */

func Test_leetcode_053(t *testing.T) {
	fmt.Println(maxSubArray([]int{-2, 1, -3, 4, -1, 2, 1, -5, 4}))
	fmt.Println(maxSubArray([]int{1}))
	fmt.Println(maxSubArray([]int{0}))
	fmt.Println(maxSubArray([]int{-1}))
	fmt.Println(maxSubArray([]int{-100000}))
}

func maxSubArray1(nums []int) int {
	defer timeCost()()
	max, length := nums[0], len(nums)
	for i := 1; i < length; i++ {
		if nums[i]+nums[i-1] > nums[i] {
			nums[i] = nums[i] + nums[i-1]
		}
		if nums[i] > max {
			max = nums[i]
		}
	}
	return max
}

func maxSubArray(nums []int) int {
	defer timeCost()()
	max, sum, length := nums[0], 0, len(nums)
	for i := 1; i < length; i++ {
		if sum > 0 {
			sum += nums[i]
		} else {
			sum = nums[i]
		}
		if max < sum {
			max = sum
		}
	}
	return max
}
