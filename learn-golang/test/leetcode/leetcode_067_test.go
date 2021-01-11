package leetcode

import (
	"fmt"
	"strconv"
	"testing"
)

/**
 * 67. 二进制求和
 */

func Test_leetcode_067(t *testing.T) {
	fmt.Println(addBinary("11", "1"))
}

func addBinary(a string, b string) string {
	defer timeCost()()
	result, aSub, bSub, carry := "", len(a)-1, len(b)-1, 0
	for aSub >= 0 || bSub >= 0 {
		if aSub >= 0 && bSub >= 0 {
			carry += int(a[aSub]) - 48 + int(b[bSub]) - 48
			result = strconv.Itoa(carry%2) + result
			carry = carry / 2
			aSub--
			bSub--
		} else if aSub >= 0 {
			carry += int(a[aSub]) - 48
			result = strconv.Itoa(carry%2) + result
			carry = carry / 2
			aSub--
		} else if bSub >= 0 {
			carry += int(b[bSub]) - 48
			result = strconv.Itoa(carry%2) + result
			carry = carry / 2
			bSub--
		}
	}
	if carry != 0 {
		return "1" + result
	}
	return result
}
