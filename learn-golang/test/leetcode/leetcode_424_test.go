package leetcode

import (
	"fmt"
	"testing"
)

/**
 * 424. 替换后的最长重复字符
 */

func Test_leetcode_424(t *testing.T) {
	fmt.Println(characterReplacement("ABAB", 2))
	fmt.Println(characterReplacement("AABABBA", 1))
}

func characterReplacement(s string, k int) int {
	defer timeCost()()
	max := func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}
	cnt := [26]int{}
	maxCnt, left := 0, 0
	for right, ch := range s {
		cnt[ch-'A']++
		maxCnt = max(maxCnt, cnt[ch-'A'])
		if right-left+1-maxCnt > k {
			cnt[s[left]-'A']--
			left++
		}
	}
	return len(s) - left
}
