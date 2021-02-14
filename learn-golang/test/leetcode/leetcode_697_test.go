package leetcode

import (
	"fmt"
	"testing"
)

/**
 * 697. 数组的度
 *
 */
func Test_leetcode_697(t *testing.T) {
	fmt.Println(findShortestSubArray([]int{1, 2, 2, 3, 1}))
	fmt.Println(findShortestSubArray([]int{1, 2, 2, 3, 1}))
	fmt.Println(findShortestSubArray([]int{1, 2, 2, 3, 1, 4, 2}))
}

func findShortestSubArray(nums []int) int {
	defer timeCost()()
	mp, result := make(map[int][]int), []int{0, 1, 0}
	for i, num := range nums {
		if _, OK := mp[num]; OK {
			list := mp[num]
			list[0]++
			list[2] = i
		} else {
			mp[num] = []int{1, i, i}
		}
	}
	for _, v := range mp {
		if result[0] < v[0] || (result[0] == v[0] && result[2]-result[1] > v[2]-v[1]) {
			result = v
		}
	}
	return result[2] - result[1] + 1
}
