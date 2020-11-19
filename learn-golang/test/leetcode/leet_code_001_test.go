package testing

import (
	"errors"
	"fmt"
	"testing"
)

/**
 * 1. 两数之和
 * 给定一个整数数组 nums 和一个目标值 target，请你在该数组中找出和为目标值的那 两个 整数，并返回他们的数组下标。
 * 你可以假设每种输入只会对应一个答案。但是，数组中同一个元素不能使用两遍。
 * 示例:
 * 给定 nums = [2, 7, 11, 15], target = 9
 * 因为 nums[0] + nums[1] = 2 + 7 = 9
 * 所以返回 [0, 1]
 */
func Test_leet_code_001(t *testing.T) {
	var num []int = []int{2, 7, 11, 15}
	var target int = 9
	result, _ := twoSum(num, target)
	fmt.Printf("%v", result)

}

func twoSum(nums []int, target int) ([]int, error) {
	var maps map[int]int = make(map[int]int)
	for i := 0; i < len(nums); i++ {
		var temp int = target - nums[i]
		if value, ok := maps[temp]; ok {
			return []int{value, i}, nil
		}
		maps[nums[i]] = i
	}
	return nil, errors.New("error")
}
