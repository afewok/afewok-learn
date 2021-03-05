package leetcode

import (
	"fmt"
	"testing"
)

/**
 * 剑指 Offer 65. 不用加减乘除做加法
 */
func Test_leetcode_offer_065(t *testing.T) {
	fmt.Println(add065(1, 1))
}

func add065(a int, b int) int {
	defer timeCost()()
	for b != 0 { // 当进位为 0 时跳出
		c := (a & b) << 1 // c = 进位
		a ^= b            // a = 非进位和
		b = c             // b = 进位
	}
	return a
}
