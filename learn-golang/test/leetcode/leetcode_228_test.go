package leetcode

import (
	"fmt"
	"strconv"
	"testing"
)

/**
 * 228. 汇总区间
 */

func Test_leetcode_5650(t *testing.T) {
	fmt.Println(summaryRanges([]int{0, 1, 2, 4, 5, 7}))
	fmt.Println(summaryRanges([]int{0, 2, 3, 4, 6, 8, 9}))
	fmt.Println(summaryRanges([]int{}))
	fmt.Println(summaryRanges([]int{-1}))
	fmt.Println(summaryRanges([]int{0}))
}

func summaryRanges(nums []int) []string {
	defer timeCost()()
	result, sub1, sub2 := make([]string, 0), -1, -1
	for i, num := range nums {
		if sub1 == -1 {
			sub1 = i
			sub2 = i
		} else if nums[sub2]+1 == num {
			sub2 = i
		} else {
			if sub1 == sub2 {
				result = append(result, strconv.Itoa(nums[sub1]))
			} else {
				result = append(result, strconv.Itoa(nums[sub1])+"->"+strconv.Itoa(nums[sub2]))
			}

			sub1 = i
			sub2 = i
		}
	}
	if sub1 > -1 {
		if sub1 == sub2 {
			result = append(result, strconv.Itoa(nums[sub1]))
		} else {
			result = append(result, strconv.Itoa(nums[sub1])+"->"+strconv.Itoa(nums[sub2]))
		}
	}

	return result
}
