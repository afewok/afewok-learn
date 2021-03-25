package leetcode

import (
	"fmt"
	"testing"
)

/**
 * 198. 打家劫舍
 */

func Test_leetcode_198(t *testing.T) {
	fmt.Println(rob([]int{1, 2, 3, 1}))
	fmt.Println(rob([]int{2, 7, 9, 3, 1}))
}

func rob1(nums []int) int {
	defer timeCost()()
	length := len(nums)
	return rob198(nums, 0, length, false)
}

func rob198(nums []int, sub, length int, OK bool) int {
	if sub == length {
		return 0
	}
	if OK {
		return rob198(nums, sub+1, length, false)
	} else {
		temp1 := nums[sub] + rob198(nums, sub+1, length, true)
		temp2 := rob198(nums, sub+1, length, false)
		if temp1 > temp2 {
			return temp1
		}
		return temp2
	}
}

func rob(nums []int) int {
	defer timeCost()()
	length := len(nums)
	if length == 0 {
		return 0
	} else if length == 1 {
		return nums[0]
	}
	max := func(x, y int) int {
		if x > y {
			return x
		}
		return y
	}
	dp := make([]int, length+1)
	dp[0] = 0
	dp[1] = nums[0]
	for i := 2; i <= length; i++ {
		dp[i] = max(dp[i-1], nums[i-1]+dp[i-2])
	}
	return dp[length]
}
