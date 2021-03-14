package leetcode

import (
	"fmt"
	"testing"
)

/**
 * 164. 最大间距
 */

func Test_leetcode_164(t *testing.T) {
	fmt.Println(maximumGap([]int{3, 6, 9, 1}))
	fmt.Println(maximumGap([]int{10}))
}

func maximumGap(nums []int) int {
	defer timeCost()()
	length, max := len(nums), 0
	for i := 0; i < length; i++ {
		for j := i + 1; j < length; j++ {
			if nums[i] > nums[j] {
				nums[i], nums[j] = nums[j], nums[i]
			}
		}
		if i > 0 && max < nums[i]-nums[i-1] {
			max = nums[i] - nums[i-1]
		}
	}
	return max
}

func maximumGap1(nums []int) int {
	n := len(nums)
	if n < 2 {
		return 0
	}
	var buckets [10][]int
	for i := 0; i < 10; i++ {
		buckets[i] = make([]int, 0)
	}
	maxVal := nums[0]
	for i := 1; i < n; i++ {
		if nums[i] > maxVal {
			maxVal = nums[i]
		}
	}
	if maxVal == 0 {
		return 0
	}
	base := 1
	for base <= maxVal {
		for i := 0; i < n; i++ {
			num := nums[i]
			buckets[(num/base)%10] = append(buckets[(num/base)%10], num)
		}
		i := 0
		for j := 0; j < 10; j++ {
			for _, v := range buckets[j] {
				nums[i] = v
				i++
			}
			buckets[j] = buckets[j][:0]
		}
		base = base * 10
	}
	maxDiff := 0
	for i := 1; i < n; i++ {
		if nums[i]-nums[i-1] > maxDiff {
			maxDiff = nums[i] - nums[i-1]
		}
	}
	return maxDiff
}
