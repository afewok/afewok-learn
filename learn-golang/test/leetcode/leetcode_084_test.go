package leetcode

import (
	"fmt"
	"testing"
)

/**
 * 84. 柱状图中最大的矩形
 */
func Test_leetcode_084(t *testing.T) {
	fmt.Println(largestRectangleArea([]int{2, 1, 5, 6, 2, 3}))
}

func largestRectangleArea(heights []int) int {
	defer timeCost()()
	f := func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}
	// 首尾添加负数高度，这样原本的第一个高度能形成升序，原本的最后一个高度也能得到处理
	heights = append([]int{-2}, heights...)
	heights = append(heights, -1)
	size := len(heights)
	// 递增栈
	s := make([]int, 1, size)

	res := 0
	i := 1
	for i < len(heights) {
		// 递增则入栈
		if heights[s[len(s)-1]] < heights[i] {
			s = append(s, i)
			i++
			continue
		}
		// s[len(s)-2]是矩形的左边界
		res = f(res, heights[s[len(s)-1]]*(i-s[len(s)-2]-1))
		s = s[:len(s)-1]
	}
	return res
}
