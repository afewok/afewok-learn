package leetcode

import (
	"fmt"
	"testing"
)

/**
 * 665. 非递减数列
 */

func Test_leetcode_665(t *testing.T) {
	fmt.Println(checkPossibility([]int{4, 2, 3}))
	fmt.Println(checkPossibility([]int{4, 2, 1}))
}

func checkPossibility(nums []int) bool {
	defer timeCost()()
	changeNum := 0
	for i := 1; i <= len(nums)-3; i++ {
		if nums[i] > nums[i+1] {
			//removable discontinuity
			if nums[i-1] <= nums[i+1] {
				changeNum++
			} else if nums[i] <= nums[i+2] {
				changeNum++
			} else {
				//jump discontinuity
				return false
			}
		}
	}
	if len(nums) >= 2 && nums[0] > nums[1] {
		changeNum++
	}
	if len(nums) >= 3 && nums[len(nums)-2] > nums[len(nums)-1] {
		changeNum++
	}
	return changeNum == 0 || changeNum == 1
}
