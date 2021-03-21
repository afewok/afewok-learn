package leetcode

import (
	"fmt"
	"testing"
)

/**
 * 5709. 最大升序子数组和
 */

func Test_leetcode_5709(t *testing.T) {
	fmt.Println(maxAscendingSum([]int{10, 20, 30, 5, 10, 50}))
	fmt.Println(maxAscendingSum([]int{10, 20, 30, 40, 50}))
	fmt.Println(maxAscendingSum([]int{12, 17, 15, 13, 10, 11, 12}))
	fmt.Println(maxAscendingSum([]int{100, 10, 1}))
}

func maxAscendingSum(nums []int) int {
	defer timeCost()()
	max, temp, length := 0, nums[0], len(nums)
	for i := 1; i < length; i++ {
		if nums[i-1] >= nums[i] {
			if max < temp {
				max = temp
			}
			temp = 0
		}
		temp += nums[i]
	}
	if max < temp {
		max = temp
	}
	return max
}
