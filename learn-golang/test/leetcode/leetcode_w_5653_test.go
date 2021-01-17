package leetcode

import (
	"fmt"
	"testing"
)

/**
 * 5653. 可以形成最大正方形的矩形数目
 */

func Test_leetcode_5653(t *testing.T) {
	fmt.Println(countGoodRectangles([][]int{{5, 8}, {3, 9}, {5, 12}, {16, 5}}))
	fmt.Println(countGoodRectangles([][]int{{2, 3}, {3, 7}, {4, 3}, {3, 7}}))
}

func countGoodRectangles(rectangles [][]int) int {
	defer timeCost()()
	mp := make(map[int]int)
	for _, rectangle := range rectangles {
		if rectangle[0] < rectangle[1] {
			mp[rectangle[0]]++
		} else {
			mp[rectangle[1]]++
		}
	}
	max := 0
	for k := range mp {
		if max < k {
			max = k
		}
	}
	return mp[max]
}
