package leetcode

import (
	"fmt"
	"testing"
)

/**
 * 342. 4的幂
 */

func Test_leetcode_342(t *testing.T) {
	fmt.Println(isPowerOfFour(16))
	fmt.Println(isPowerOfFour(5))
	fmt.Println(isPowerOfFour(1))
}

func isPowerOfFour(n int) bool {
	defer timeCost()()
	for n >= 1 {
		if n == 1 {
			return true
		}
		if n%4 != 0 {
			return false
		}
		n /= 4
	}
	return false
}
