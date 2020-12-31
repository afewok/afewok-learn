package leetcode

import (
	"fmt"
	"sort"
	"testing"
)

/**
 * 435. 无重叠区间=
 *
 */
func Test_leetcode_435(t *testing.T) {
	fmt.Println(eraseOverlapIntervals([][]int{{1, 2}, {2, 3}, {3, 4}, {1, 3}}))
}

func eraseOverlapIntervals(intervals [][]int) int {
	defer timeCost()()
	sort.Slice(intervals, func(i, j int) bool {
		a, b := intervals[i], intervals[j]
		return a[0] < b[0] || a[0] == b[0] && a[1] < b[1]
	})
	miR := int(-1e18)
	for _, p := range intervals {
		l, r := p[0], p[1]
		if l >= miR {
			miR = r
		} else {
			ans++
			if r < miR {
				miR = r
			}
		}
	}
	return
}
