package leetcode

import (
	"fmt"
	"testing"
)

/**
 * 66. 加一
 */

func Test_leetcode_66(t *testing.T) {
	fmt.Println(plusOne([]int{1, 2, 3}))
	fmt.Println(plusOne([]int{9, 9, 9}))
}

func plusOne(digits []int) []int {
	defer timeCost()()
	carry, length := 1, len(digits)
	for i := length - 1; i >= 0; i-- {
		if digits[i]+carry >= 10 {
			digits[i] = (digits[i] + carry) - 10
		} else {
			digits[i] += carry
			carry = 0
		}
	}
	if carry > 0 {
		return append([]int{1}, digits...)
	}
	return digits
}
