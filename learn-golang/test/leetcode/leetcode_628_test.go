package leetcode

import (
	"fmt"
	"sort"
	"testing"
)

/**
 * 628. 三个数的最大乘积
 */

func Test_leetcode_628(t *testing.T) {
	fmt.Println(removeElement([]int{3, 2, 2, 3}, 3))
	fmt.Println(removeElement([]int{0, 1, 2, 2, 3, 0, 4, 2}, 2))
}

func maximumProduct(nums []int) int {
	defer timeCost()()
	max := func(a ...int) int {
		var m = a[0]
		for _, v := range a {
			if v > m {
				m = v
			}
		}
		return m
	}
	sort.Ints(nums)
	// 情况1:正数、正数、正数
	// 情况2:负数、正数、正数
	// 情况3:负数、负数、正数
	// 情况4:负数、负数、负数
	// 情况5:含有0的情况
	return max(nums[len(nums)-1]*nums[len(nums)-2]*nums[len(nums)-3], nums[0]*nums[1]*nums[len(nums)-1])
}
