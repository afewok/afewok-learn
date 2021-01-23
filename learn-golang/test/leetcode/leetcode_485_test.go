package leetcode

import (
	"fmt"
	"testing"
)

/**
 * 485. 最大连续1的个数
 */

func Test_leetcode_485(t *testing.T) {
	fmt.Println(findMaxConsecutiveOnes([]int{0}))
	fmt.Println(findMaxConsecutiveOnes([]int{1, 1, 0, 1, 1, 1}))
}

func findMaxConsecutiveOnes(nums []int) int {
	defer timeCost()()
	max, sub := 0, -1
	for i, num := range nums {
		if num == 1 {
			if sub == -1 {
				sub = i
			}
		} else {
			if sub != -1 {
				temp := i - sub
				sub = -1
				if max < temp {
					max = temp
				}
			}
		}
	}
	if sub != -1 {
		temp := len(nums) - sub
		if max < temp {
			max = temp
		}
	}
	return max
}
