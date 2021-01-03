package leetcode

import (
	"fmt"
	"testing"
)

/**
 * 371. 两整数之和
 */

func Test_leetcode_371(t *testing.T) {
	fmt.Println(getSum(1, 2))
	fmt.Println(getSum(-2, 3))
}

func getSum(a int, b int) int {
	defer timeCost()()
	for b != 0 {
		sum := a ^ b
		carry := (a & b) << 1
		a = sum
		b = carry
	}
	return a
}
