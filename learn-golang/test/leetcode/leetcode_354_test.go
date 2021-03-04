package leetcode

import (
	"fmt"
	"sort"
	"testing"
)

/**
 * 354. 俄罗斯套娃信封问题
 */
func Test_leetcode_354(t *testing.T) {
	fmt.Println([]int{2, 2, 1})
	fmt.Println([]int{4, 1, 2, 1, 2})
}

func maxEnvelopes(envelopes [][]int) int {
	defer timeCost()()
	sort.Slice(envelopes, func(i, j int) bool {
		a, b := envelopes[i], envelopes[j]
		return a[0] < b[0] || a[0] == b[0] && a[1] > b[1]
	})

	f := []int{}
	for _, e := range envelopes {
		h := e[1]
		if i := sort.SearchInts(f, h); i < len(f) {
			f[i] = h
		} else {
			f = append(f, h)
		}
	}
	return len(f)
}
