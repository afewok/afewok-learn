package leetcode

import (
	"fmt"
	"testing"
)

/**
 * 70. 爬楼梯
 */

func Test_leetcode_070(t *testing.T) {
	fmt.Println(climbStairs(1)) //1
	fmt.Println(climbStairs(2)) //2
	fmt.Println(climbStairs(3)) //3
	fmt.Println(climbStairs(4)) //5
	fmt.Println(climbStairs(5)) //8
	fmt.Println(climbStairs(6)) //13
	fmt.Println(climbStairs(7)) //21
}

func climbStairs(n int) int {
	defer timeCost()()
	a, b := 0, 1
	for n > 1 {
		a, b = b, a+b
		n--
	}
	return a + b
}
