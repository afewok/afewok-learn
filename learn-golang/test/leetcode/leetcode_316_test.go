package leetcode

import (
	"fmt"
	"testing"
)

/**
 * 316. 去除重复字母
 */

func Test_leetcode_316(t *testing.T) {
	fmt.Println(removeDuplicateLetters("bcabc"))
}

func removeDuplicateLetters(s string) string {
	defer timeCost()()
}
