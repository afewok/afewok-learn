package leetcode

import (
	"fmt"
	"testing"
)

/**
 * 674. 最长连续递增序列
 */

func Test_leetcode_674(t *testing.T) {
	fmt.Println(findLengthOfLCIS([]int{1, 3, 5, 4, 7}))
	fmt.Println(findLengthOfLCIS([]int{2, 2, 2, 2, 2}))
}

func findLengthOfLCIS(nums []int) int {
	defer timeCost()()
	length := len(nums)
	if length < 1 {
		return 0
	}
	max, sub := 1, 0
	for i := 1; i < length; i++ {
		if nums[i-1] >= nums[i] {
			temp := i - sub
			sub = i
			if max < temp {
				max = temp
			}
		}
	}
	temp := length - sub
	if max < temp {
		max = temp
	}
	return max
}
