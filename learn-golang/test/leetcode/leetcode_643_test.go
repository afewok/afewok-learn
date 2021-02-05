package leetcode

import (
	"fmt"
	"testing"
)

/**
 * 643. 子数组最大平均数 I
 */

func Test_leetcode_643(t *testing.T) {
	fmt.Println(findMaxAverage([]int{1, 12, -5, -6, 50, 3}, 4))
}

func findMaxAverage(nums []int, k int) float64 {
	defer timeCost()()
	length, max, temp := len(nums), 0, 0
	for i := 0; i < k; i++ {
		max += nums[i]
	}
	temp = max
	for i := k; i < length; i++ {
		temp = temp - nums[i-k] + nums[i]
		if max < temp {
			max = temp
		}
	}
	return float64(max) / float64(k)
}
