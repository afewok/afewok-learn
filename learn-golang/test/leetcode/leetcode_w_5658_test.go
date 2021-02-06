package leetcode

import (
	"fmt"
	"testing"
)

/**
 * 5658. 任意子数组和的绝对值的最大值
 */

func Test_leetcode_5658(t *testing.T) {
	fmt.Println(sumOfUnique([]int{1, 2, 3, 2}))
	fmt.Println(sumOfUnique([]int{1, 1, 1, 1, 1}))
	fmt.Println(sumOfUnique([]int{1, 2, 3, 4, 5}))
}

func maxAbsoluteSum(nums []int) int {
	defer timeCost()()
	return 1

}
