package leetcode

import (
	"fmt"
	"testing"
)

/**
 * 5645. 找到最高海拔
 */

func Test_leetcode_5645(t *testing.T) {
	fmt.Println(largestAltitude([]int{-5, 1, 5, 0, -7}))
	fmt.Println(largestAltitude([]int{-4, -3, -2, -1, 4, 3, 2}))
}

func largestAltitude(gain []int) int {
	defer timeCost()()
	max, temp := 0, 0
	gain = append(gain, 0)
	for _, v := range gain {
		temp += v
		if max < temp {
			max = temp
		}
	}
	return max
}
