package leetcode

import (
	"fmt"
	"testing"
)

/**
 * 1512. 好数对的数目
 */

func Test_leetcode_1512(t *testing.T) {
	fmt.Println(numIdenticalPairs([]int{1, 2, 3, 1, 1, 3}))
	fmt.Println(numIdenticalPairs([]int{1, 1, 1, 1}))
	fmt.Println(numIdenticalPairs([]int{1, 2, 3}))
}

func numIdenticalPairs(nums []int) int {
	defer timeCost()()
	length, result := len(nums), 0
	for i := 0; i < length; i++ {
		for j := i + 1; j < length; j++ {
			if nums[i] == nums[j] {
				result++
			}
		}
	}
	return result
}
