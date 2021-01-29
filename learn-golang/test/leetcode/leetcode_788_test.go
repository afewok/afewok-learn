package leetcode

import (
	"fmt"
	"strings"
	"testing"
)

/**
 * 796. 旋转字符串
 *
 */
func Test_leetcode_788(t *testing.T) {
	fmt.Println(rotateString("abcde", "cdeab"))
	fmt.Println(rotateString("abcde", "abced"))
}

func rotateString(A string, B string) bool {
	defer timeCost()()
	lenA, lenB := len(A), len(B)
	if lenA != lenB {
		return false
	} else if lenA == 0 {
		return true
	}
	for i := 0; i < lenA; i++ {
		temp := strings.Index(B, A[i:lenA])
		if temp == 0 && strings.LastIndex(B, A[0:i]) == lenB-i {
			return true
		}
	}
	return false
}
