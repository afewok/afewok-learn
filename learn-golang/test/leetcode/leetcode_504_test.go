package leetcode

import (
	"fmt"
	"strconv"
	"testing"
)

/**
 * 504. 七进制数
 */

func Test_leetcode_504(t *testing.T) {
	fmt.Println(convertToBase7(100))
	fmt.Println(convertToBase7(-7))
}

func convertToBase7(num int) string {
	defer timeCost()()
	head, result := "", ""
	if num < 0 {
		head = "-"
		num = -1 * num
	}
	for num >= 7 {
		result = strconv.Itoa(num%7) + result
		num = num / 7
	}
	return head + strconv.Itoa(num%7) + result
}
