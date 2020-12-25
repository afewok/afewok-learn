package leetcode

import (
	"fmt"
	"testing"
)

/**
 * 168. Excel表列名称
 */

func Test_leetcode_168(t *testing.T) {
	fmt.Println(convertToTitle(1))
	fmt.Println(convertToTitle(28))
	fmt.Println(convertToTitle(52))
	fmt.Println(convertToTitle(701))
}

func convertToTitle(n int) string {
	defer timeCost()()
	table := "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	str := ""
	for n > 26 {
		temp := n % 26
		if temp == 0 {
			str = string(table[25]) + str
			n = n/26 - 1
		} else {
			str = string(table[temp-1]) + str
			n = n / 26
		}

	}
	return string(table[n-1]) + str
}
