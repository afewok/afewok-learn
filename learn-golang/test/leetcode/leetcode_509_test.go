package leetcode

import (
	"fmt"
	"testing"
)

/**
 * 509. 斐波那契数
 */

func Test_leetcode_509(t *testing.T) {
	fmt.Println(fib(2))
	fmt.Println(fib(3))
	fmt.Println(fib(4))
}

func fib(n int) int {
	defer timeCost()()
	if n == 0 {
		return 0
	} else if n == 1 || n == 2 {
		return 1
	}
	a, b := 1, 1
	for ; n >= 3; n-- {
		a, b = b, a+b
	}
	return b
}
