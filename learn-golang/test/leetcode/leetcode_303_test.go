package leetcode

import (
	"fmt"
	"testing"
)

/**
 * 303. 区域和检索 - 数组不可变
 */

func Test_leetcode_303(t *testing.T) {
	fmt.Println(getRow(3))
}

type NumArray struct {
	sum []int
}

func Constructor303(nums []int) NumArray {
	s := make([]int, len(nums)+1)
	for i := 0; i < len(nums); i++ {
		s[i+1] = s[i] + nums[i]
	}
	return NumArray{s}
}

func (this *NumArray) SumRange(i int, j int) int {
	return this.sum[j+1] - this.sum[i]
}
