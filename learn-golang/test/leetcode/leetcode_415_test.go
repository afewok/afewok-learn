package leetcode

import (
	"fmt"
	"strconv"
	"testing"
)

/**
 * 415. 字符串相加
 */

func Test_leetcode_415(t *testing.T) {
	fmt.Println(addStrings("123", "321"))
	fmt.Println(addStrings("999", "999"))
}

func addStrings(num1 string, num2 string) string {
	defer timeCost()()
	sub1, sub2, result, temp, carry := len(num1)-1, len(num2)-1, "", 0, 0
	for sub1 >= 0 || sub2 >= 0 {
		if sub1 >= 0 && sub2 >= 0 {
			temp = int(num1[sub1]) - 48 + int(num2[sub2]) - 48 + carry
			carry = temp / 10
			result = strconv.Itoa(temp%10) + result
		} else if sub1 >= 0 {
			temp = int(num1[sub1]) - 48 + carry
			carry = temp / 10
			result = strconv.Itoa(temp%10) + result
		} else if sub2 >= 0 {
			temp = int(num2[sub2]) - 48 + carry
			carry = temp / 10
			result = strconv.Itoa(temp%10) + result
		}
		sub1--
		sub2--
	}
	if carry > 0 {
		return "1" + result
	}
	return result
}
