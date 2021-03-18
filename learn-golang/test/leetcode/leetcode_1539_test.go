package leetcode

import (
	"fmt"
	"testing"
)

/**
 * 1539. 第 k 个缺失的正整数
 */

func Test_leetcode_1539(t *testing.T) {
	fmt.Println(findKthPositive([]int{2, 3, 4, 7, 11}, 5))
	fmt.Println(findKthPositive([]int{1, 2, 3, 4}, 2))
}

func findKthPositive(arr []int, k int) int {
	defer timeCost()()
	for i := len(arr) - 1; i >= 0; i-- {
		if arr[i]-(i+1) < k {
			return arr[i] + k - (arr[i] - (i + 1))
		}
	}
	return k
}
