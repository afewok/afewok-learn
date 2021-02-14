package leetcode

import (
	"fmt"
	"testing"
)

/**
 * 152. 乘积最大子数组
 */

func Test_leetcode_152(t *testing.T) {
	fmt.Println(maxProduct([]int{0, 2}))
	fmt.Println(maxProduct([]int{-2}))
	fmt.Println(maxProduct([]int{2, 3, -2, 4}))
	fmt.Println(maxProduct([]int{-2, 0, -1}))
}

func maxProduct(nums []int) int {
	defer timeCost()()
	max := func(x, y int) int {
		if x > y {
			return x
		}
		return y
	}

	min := func(x, y int) int {
		if x < y {
			return x
		}
		return y
	}
	maxF, minF, ans := nums[0], nums[0], nums[0]
	for i := 1; i < len(nums); i++ {
		mx, mn := maxF, minF
		maxF = max(mx*nums[i], max(nums[i], mn*nums[i]))
		minF = min(mn*nums[i], min(nums[i], mx*nums[i]))
		ans = max(maxF, ans)
	}
	return ans
}
