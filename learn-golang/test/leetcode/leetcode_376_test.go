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
 *
 * 排除连续上升，或者连续下降，以及相等的情况就可以了
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
			up[i] = maxTwo(up[i-1], down[i-1]+1)
			down[i] = down[i-1]
		} else if nums[i] < nums[i-1] {
			up[i] = up[i-1]
			down[i] = maxTwo(up[i-1]+1, down[i-1])
		} else {
			up[i] = up[i-1]
			down[i] = down[i-1]
		}
	}
	return maxTwo(up[length-1], down[length-1])
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
			up = maxTwo(up, down+1)
		} else if nums[i] < nums[i-1] {
			down = maxTwo(up+1, down)
		}
	}
	return maxTwo(up, down)
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
	return maxTwo(up, down)
}

func wiggleMaxLength4(nums []int) int {
	defer timeCost()()
	length := len(nums)
	if length < 2 {
		return length
	}
	result := 1
	preDiff := nums[1] - nums[0]
	if preDiff != 0 {
		result = 2
	}
	for i := 2; i < length; i++ {
		diff := nums[i] - nums[i-1]
		if (diff > 0 && preDiff <= 0) || (diff < 0 && preDiff >= 0) {
			result++
			preDiff = diff
		}
	}
	return result
}

func wiggleMaxLength(nums []int) int {
	defer timeCost()()
	length := len(nums)
	if length < 2 {
		return length
	}
	i, preDiff, diff, result := 0, 0, 0, 1
	for i = 1; i < length; i++ {
		diff = nums[i] - nums[i-1]
		if diff != 0 {
			result++
			preDiff = diff
			break
		}
	}
	for i = i + 1; i < length; i++ {
		diff = nums[i] - nums[i-1]
		if (diff > 0 && preDiff < 0) || (diff < 0 && preDiff > 0) {
			result++
			preDiff = diff
		}
	}
	return result
}
