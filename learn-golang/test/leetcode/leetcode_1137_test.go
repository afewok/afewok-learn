package leetcode

import (
	"fmt"
	"testing"
)

/**
 * 1137. 第 N 个泰波那契数
 */

func Test_leetcode_1137(t *testing.T) {
	fmt.Println(tribonacci(4))
	fmt.Println(tribonacci(25))
}

func tribonacci(n int) int {
	defer timeCost()()
	t1, t2, t3 := 0, 1, 1
	if n == 0 {
		return t1
	} else if n == 1 {
		return t2
	} else if n == 2 {
		return t3
	}
	for i := 3; i <= n; i++ {
		t1, t2, t3 = t2, t3, t1+t2+t3
	}
	return t3
}
