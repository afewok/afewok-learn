package leetcode

import (
	"fmt"
	"sort"
	"testing"
)

/**
 * 561. 数组拆分 I
 */

func Test_leetcode_561(t *testing.T) {
	fmt.Println(arrayPairSum([]int{1, 4, 3, 2}))
	fmt.Println(arrayPairSum([]int{6, 2, 6, 5, 1, 2}))
}

func arrayPairSum(nums []int) int {
	defer timeCost()()
	sort.Sort(sort.IntSlice(nums))
	length, result := len(nums), 0
	for i := 0; i < length; i += 2 {
		result += nums[i]
	}
	return result
}
