package leetcode

import (
	"fmt"
	"testing"
)

/**
 * 567. 字符串的排列
 */

func Test_leetcode_567(t *testing.T) {
	fmt.Println(checkInclusion("ab", "eidbaooo"))
	fmt.Println(checkInclusion("ab", "eidboaoo"))
}

func checkInclusion(s1 string, s2 string) bool {
	defer timeCost()()
	return false
}
