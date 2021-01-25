package leetcode

import (
	"fmt"
	"math"
	"testing"
)

/**
 * 453. 最小操作次数使数组元素相等
 */

func Test_leetcode_453(t *testing.T) {
	fmt.Println(minMoves([]int{1, 2, 3}))
}

func minMoves(nums []int) int {
	defer timeCost()()
	min, count := math.MaxInt32, 0
	for _, num := range nums {
		if min > num {
			min = num
		}
	}
	for _, num := range nums {
		count += num - min
	}
	return count
}
