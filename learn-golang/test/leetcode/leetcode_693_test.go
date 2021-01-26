package leetcode

import (
	"fmt"
	"testing"
)

/**
 * 693. 交替位二进制数
 *
 */
func Test_leetcode_693(t *testing.T) {
	fmt.Println(hasAlternatingBits(5))
	fmt.Println(hasAlternatingBits(7))
	fmt.Println(hasAlternatingBits(11))
	fmt.Println(hasAlternatingBits(10))
	fmt.Println(hasAlternatingBits(3))
}

func hasAlternatingBits(n int) bool {
	defer timeCost()()
	c := n % 2
	n = n / 2
	for n > 0 {
		temp := n % 2
		if c == temp {
			return false
		}
		c = temp
		n = n / 2
	}
	return true
}
