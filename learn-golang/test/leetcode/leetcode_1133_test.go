package leetcode

import (
	"fmt"
	"testing"
)

/**
 * 1133. 最大唯一数
 */

func Test_leetcode_1133(t *testing.T) {
	fmt.Println(largestUniqueNumber([]int{5, 7, 3, 9, 4, 9, 8, 3, 1}))
	fmt.Println(largestUniqueNumber([]int{9, 9, 8, 8}))
}

func largestUniqueNumber(A []int) int {
	defer timeCost()()
	max := -1 << 10
	mp := make(map[int]int)
	for _, num := range A {
		mp[num]++
	}

	for k, v := range mp {
		if v == 1 && max < k {
			max = k
		}
	}
	if max == -1<<10 {
		return -1
	}
	return max
}
