package leetcode

import (
	"fmt"
	"testing"
)

/**
 * 976. 三角形的最大周长
 */
func Test_leetcode_976(t *testing.T) {
	// fmt.Println(largestPerimeter([]int{2, 1, 2}))
	// fmt.Println(largestPerimeter([]int{1, 2, 1}))
	fmt.Println(largestPerimeter([]int{3, 2, 3, 4}))
	fmt.Println(largestPerimeter([]int{3, 6, 2, 3}))
}

func largestPerimeter(nums []int) int {
	defer timeCost()()
	length := len(nums)
	for i := 0; i < length; i++ {
		for j := i + 1; j < length; j++ {
			if nums[i] < nums[j] {
				nums[i], nums[j] = nums[j], nums[i]
			}
		}
	}
	for i := 2; i < length; i++ {
		for i1 := 0; i1 <= i; i1++ {
			for i2 := i1 + 1; i2 <= i; i2++ {
				for i3 := i2 + 1; i3 <= i; i3++ {
					if i1 == i2 || i1 == i3 || i2 == i3 {
						continue
					}
					if nums[i1] >= nums[i2]+nums[i3] || nums[i2] >= nums[i1]+nums[i3] || nums[i3] >= nums[i1]+nums[i2] {
						continue
					}
					return nums[i1] + nums[i2] + nums[i3]
				}
			}
		}
	}

	return 0
}
