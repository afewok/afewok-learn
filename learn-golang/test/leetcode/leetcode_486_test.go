package leetcode

import (
	"fmt"
	"testing"
)

/**
 * 486. 预测赢家
 */

func Test_leetcode_486(t *testing.T) {
	fmt.Println(PredictTheWinner([]int{1, 5, 2}))
	fmt.Println(PredictTheWinner([]int{1, 5, 233, 7}))
}

func PredictTheWinner(nums []int) bool {
	defer timeCost()()
	max := func(x, y int) int {
		if x > y {
			return x
		}
		return y
	}
	length := len(nums)
	dp := make([]int, length)
	for i := 0; i < length; i++ {
		dp[i] = nums[i]
	}
	for i := length - 2; i >= 0; i-- {
		for j := i + 1; j < length; j++ {
			dp[j] = max(nums[i]-dp[j], nums[j]-dp[j-1])
		}
	}
	return dp[length-1] >= 0
}
