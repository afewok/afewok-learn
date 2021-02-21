package leetcode

import (
	"fmt"
	"testing"
)

/**
 * 1438. 绝对差不超过限制的最长连续子数组
 */
func Test_leetcode_1438(t *testing.T) {
	fmt.Println(longestSubarray([]int{2, 5, 2}, 9))
	fmt.Println(longestSubarray([]int{1, 5, 6, 7, 8, 10, 6, 5, 6}, 4))
	fmt.Println(longestSubarray([]int{8, 2, 4, 7}, 4))
	fmt.Println(longestSubarray([]int{10, 1, 2, 4, 7, 2}, 5))
	fmt.Println(longestSubarray([]int{4, 2, 2, 2, 4, 4, 2, 2}, 0))
}

func longestSubarray(nums []int, limit int) int {
	defer timeCost()()
	sub, minSub, maxSub, result := 0, 0, 0, 0
	for i, num := range nums {
		if nums[minSub] > num {
			minSub = i
			if nums[maxSub]-num > limit {
				sub, maxSub = maxSub+1, i
				for j := i; j >= sub; j-- {
					if nums[j]-num > limit {
						sub = j + 1
						break
					}
					if nums[maxSub] < nums[j] {
						maxSub = j
					}
				}
			}
		} else if nums[maxSub] < num {
			maxSub = i
			if num-nums[minSub] > limit {
				sub, minSub = minSub+1, i
				for j := i; j >= sub; j-- {
					if num-nums[j] > limit {
						sub = j + 1
						break
					}
					if nums[minSub] > nums[j] {
						minSub = j
					}
				}
			}

		}
		if result < i-sub+1 {
			result = i - sub + 1
		}
	}
	return result
}
