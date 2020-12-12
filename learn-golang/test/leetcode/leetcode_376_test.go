package leetcode

import (
	"fmt"
	"testing"
)

/**
 * 376. 摆动序列
 *
 * 思路：动态规划、贪心
 * DP算法（Dynamic Programming,俗称动态规划）
 */
func Test_leetcode_376(t *testing.T) {
	fmt.Println(wiggleMaxLength([]int{1, 7, 4, 9, 2, 5}))
	fmt.Println(wiggleMaxLength([]int{1, 17, 5, 10, 13, 15, 10, 5, 16, 8}))
	fmt.Println(wiggleMaxLength([]int{1, 2, 3, 4, 5, 6, 7, 8, 9}))
}

func wiggleMaxLength1(nums []int) int {
	defer timeCost()()
	length := len(nums)
	if length < 2 {
		return length
	}
	up := make([]int, length)
	down := make([]int, length)
	up[0] = 1
	down[0] = 1
	for i := 1; i < length; i++ {
		if nums[i] > nums[i-1] {
			up[i] = max(up[i-1], down[i-1]+1)
			down[i] = down[i-1]
		} else if nums[i] < nums[i-1] {
			up[i] = up[i-1]
			down[i] = max(up[i-1]+1, down[i-1])
		} else {
			up[i] = up[i-1]
			down[i] = down[i-1]
		}
	}
	return max(up[length-1], down[length-1])
}

func wiggleMaxLength2(nums []int) int {
	defer timeCost()()
	length := len(nums)
	if length < 2 {
		return length
	}
	up, down := 1, 1
	for i := 1; i < length; i++ {
		if nums[i] > nums[i-1] {
			up = max(up, down+1)
		} else if nums[i] < nums[i-1] {
			down = max(up+1, down)
		}
	}
	return max(up, down)
}

func wiggleMaxLength3(nums []int) int {
	defer timeCost()()
	length := len(nums)
	if length < 2 {
		return length
	}
	up, down := 1, 1
	for i := 1; i < length; i++ {
		if nums[i] > nums[i-1] {
			up = down + 1
		} else if nums[i] < nums[i-1] {
			down = up + 1
		}
	}
	return max(up, down)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func wiggleMaxLength(nums []int) int {
	defer timeCost()()
	length := len(nums)
	if length < 2 {
		return length
	}
	ans := 1
	prevDiff := nums[1] - nums[0]
	if prevDiff != 0 {
		ans = 2
	}
	for i := 2; i < length; i++ {
		diff := nums[i] - nums[i-1]
		if (diff > 0 && prevDiff <= 0) || (diff < 0 && prevDiff >= 0) {
			ans++
			prevDiff = diff
		}
	}
	return ans
}
