package leetcode

import (
	"fmt"
	"testing"
)

/**
 * 326. 3的幂
 */

func Test_leetcode_326(t *testing.T) {
	fmt.Println(isPowerOfThree(27))
	fmt.Println(isPowerOfThree(0))
	fmt.Println(isPowerOfThree(9))
	fmt.Println(isPowerOfThree(45))
	fmt.Println(isPowerOfThree(1))
}

func isPowerOfThree(n int) bool {
	defer timeCost()()
	for n >= 1 {
		if n == 1 {
			return true
		}
		if n%3 != 0 {
			return false
		}
		n /= 3
	}
	return false
}
