package leetcode

import (
	"fmt"
	"sort"
	"testing"
)

/**
 * 15. 三数之和
 */

func Test_leetcode_015(t *testing.T) {
	fmt.Println(removeElement([]int{3, 2, 2, 3}, 3))
	fmt.Println(removeElement([]int{0, 1, 2, 2, 3, 0, 4, 2}, 2))
}

func threeSum(nums []int) [][]int {
	defer timeCost()()
	sort.Ints(nums)
	var saveArray [][]int
	var addnum int
	//index为中间值
	var index, start, end int
	for index = 1; index < len(nums)-1; index++ {
		start, end = 0, len(nums)-1
		//去重的关键步骤
		if index > 1 && nums[index] == nums[index-1] {
			start = index - 1
		}
		for start < index && index < end {
			addnum = nums[start] + nums[index] + nums[end]
			if start > 0 && nums[start] == nums[start-1] {
				start++
				continue
			}
			if end < len(nums)-1 && nums[end] == nums[end+1] {
				end--
				continue
			}
			if addnum == 0 {
				saveArray = append(saveArray, []int{nums[start], nums[index], nums[end]})
				start++
				end--
			} else if addnum > 0 {
				end--
			} else {
				start++
			}
		}
	}
	return saveArray
}
