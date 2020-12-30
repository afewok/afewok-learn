package leetcode

import (
	"fmt"
	"testing"
)

/**
 * 231. 2的幂
 */

func Test_leetcode_231(t *testing.T) {
	fmt.Println(isPowerOfTwo(1))
	fmt.Println(isPowerOfTwo(16))
	fmt.Println(isPowerOfTwo(218))
}

func isPowerOfTwo(n int) bool {
	defer timeCost()()
	for n > 0 {
		if n&1 == 1 {
			return n == 1
		}
		n = n >> 1
	}
	return false
}
