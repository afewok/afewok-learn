package leetcode

import (
	"fmt"
	"testing"
)

/**
 * 5697. 检查二进制字符串字段
 */

func Test_leetcode_5697(t *testing.T) {
	fmt.Println(checkOnesSegment("1001"))
	fmt.Println(checkOnesSegment("110"))
}

func checkOnesSegment(s string) bool {
	defer timeCost()()
	length := len(s)
	OK := true
	for i := 1; i < length; i++ {
		if s[i] == '1' {
			if !OK {
				return false
			}
		} else {
			OK = false
		}
	}
	return true
}
