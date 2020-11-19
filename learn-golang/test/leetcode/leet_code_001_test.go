package testing

import (
	"fmt"
	"testing"
)

/**
 * 1. 两数之和
 * <p>
 * 给定一个整数数组 nums 和一个目标值 target，请你在该数组中找出和为目标值的那 两个 整数，并返回他们的数组下标。
 * 你可以假设每种输入只会对应一个答案。但是，数组中同一个元素不能使用两遍。
 * <p>
 * 示例:
 * 给定 nums = [2, 7, 11, 15], target = 9
 * 因为 nums[0] + nums[1] = 2 + 7 = 9
 * 所以返回 [0, 1]
 *
 * 思路：暴力2次循环、头尾递进、hash字典、数组存储差值
 */
func Test_leet_code_001(t *testing.T) {
	var num []int = []int{2, 7, 11, 15}
	var target int = 9
	result1 := twoSum1(num, target)
	fmt.Printf("%v", result1)

	result2 := twoSum2(num, target)
	fmt.Printf("%v", result2)
}

func twoSum1(nums []int, target int) []int {
	var temps []int = make([]int, len(nums))
	for i := 0; i < len(nums); i++ {
		for j := 0; j < i; j++ {
			if nums[i] == temps[j] {
				return []int{j, i}
			}
		}
		temps[i] = target - nums[i]
	}
	return nil
}

func twoSum2(nums []int, target int) []int {
	var maps map[int]int = make(map[int]int, len(nums))
	for i := 0; i < len(nums); i++ {
		if value, ok := maps[nums[i]]; ok {
			return []int{value, i}
		}
		maps[target-nums[i]] = i
	}
	return nil
}
