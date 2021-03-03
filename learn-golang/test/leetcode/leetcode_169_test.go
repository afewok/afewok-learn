package leetcode

import (
	"fmt"
	"testing"
)

/**
 * 169. 多数元素
 */

func Test_leetcode_169(t *testing.T) {
	fmt.Println(majorityElement([]int{3, 2, 3}))
	fmt.Println(majorityElement([]int{2, 2, 1, 1, 1, 2, 2}))
	fmt.Println(majorityElement([]int{4, 5, 4}))
}

func majorityElement(nums []int) int {
	defer timeCost()()
	mp := make(map[int]int)
	for _, num := range nums {
		mp[num]++
	}
	for k, v := range mp {
		if v > len(nums)/2 {
			return k
		}
	}
	return 0
}

//排序
func majorityElement2(nums []int) int {
	defer timeCost()()
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[i] > nums[j] {
				nums[i], nums[j] = nums[j], nums[i]
			}
		}
	}
	return nums[len(nums)/2]
}

//分治
// func majorityElement(nums []int) int {
// 	defer timeCost()()

// }
