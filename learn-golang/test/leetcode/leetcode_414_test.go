package leetcode

import (
	"fmt"
	"testing"
)

/**
 * 414. 第三大的数
 */

func Test_leetcode_414(t *testing.T) {
	fmt.Println(thirdMax([]int{3, 2, 1}))
	fmt.Println(thirdMax([]int{1, 2}))
	fmt.Println(thirdMax([]int{2, 2, 3, 1}))
}

func thirdMax(nums []int) int {
	defer timeCost()()
	max1, max2, max3 := ^int(^uint(0)>>1), ^int(^uint(0)>>1), ^int(^uint(0)>>1)
	for _, num := range nums {
		if num > max1 {
			max3, max2, max1 = max2, max1, num
		} else if num > max2 && num < max1 {
			max3, max2 = max2, num
		} else if num > max3 && num < max2 {
			max3 = num
		}
	}
	if max3 == ^int(^uint(0)>>1) {
		return max1
	}
	return max3
}
