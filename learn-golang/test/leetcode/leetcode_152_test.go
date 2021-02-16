package leetcode

import (
	"fmt"
	"testing"
)

/**
 * 152. 乘积最大子数组
 */

func Test_leetcode_152(t *testing.T) {
	fmt.Println(maxProduct([]int{-2, -3, -4}))
	fmt.Println(maxProduct([]int{-2}))
	fmt.Println(maxProduct([]int{2, 3, -2, 4}))
	fmt.Println(maxProduct([]int{-2, 0, -1}))
}

func maxProduct(nums []int) int {
	defer timeCost()()
	max := func(x, y, z int) int {
		if x >= y && x >= z {
			return x
		} else if y >= x && y >= z {
			return y
		}
		return z
	}

	min := func(x, y, z int) int {
		if x <= y && x <= z {
			return x
		} else if y <= x && y <= z {
			return y
		}
		return z
	}
	maxF, minF, ans := nums[0], nums[0], nums[0]
	for i := 1; i < len(nums); i++ {
		mx, mn := maxF, minF
		maxF = max(mx*nums[i], nums[i], mn*nums[i])
		minF = min(mn*nums[i], nums[i], mx*nums[i])
		ans = max(maxF, minF, ans)
	}
	return ans
}
