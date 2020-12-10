package leetcode

import (
	"fmt"
	"testing"
)

/**
 * 28. 实现 strStr()
 *
 * 思路：首字母比对、KMP 算法（Knuth-Morris-Pratt 算法）
 */
func Test_leetcode_028(t *testing.T) {
	fmt.Println(strStr("hello", "ll"))
	fmt.Println(strStr("aaaaa", "bba"))
}

func strStr(haystack string, needle string) int {
	hLen, nLen := len(haystack), len(needle)
	if nLen == 0 {
		return 0
	} else if hLen == 0 || hLen < nLen {
		return -1
	}

	for i := 0; i <= hLen-nLen; i++ {
		if haystack[i:i+nLen] == needle {
			return i
		}
	}
	return -1
}
