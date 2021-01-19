package leetcode

import (
	"fmt"
	"testing"
)

/**
 * 448. 找到所有数组中消失的数字
 */

func Test_leetcode_448(t *testing.T) {
	fmt.Println(findDisappearedNumbers([]int{4, 3, 2, 7, 8, 2, 3, 1}))
}

func findDisappearedNumbers(nums []int) []int {
	defer timeCost()()
	result := make([]int, 0, 0)
	for _, value := range nums {
		if value < 0 {
			value = -value
		}
		if nums[value-1] > 0 {
			nums[value-1] = -nums[value-1]
		}
	}
	for index, value := range nums {
		if value > 0 {
			result = append(result, index+1)
		}
	}
	return result
}
