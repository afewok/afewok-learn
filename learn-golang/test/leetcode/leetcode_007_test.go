package leetcode

import (
	"fmt"
	"testing"
)

/**
 * 7. 整数反转
 */
func Test_leetcode_007(t *testing.T) {
	fmt.Println(reverse(123))
	fmt.Println(reverse(-123))
	fmt.Println(reverse(120))
	fmt.Println(reverse(0))
	fmt.Println(reverse(1563847412))
	fmt.Println(reverse(-1563847412))
}

func reverse(x int) int {
	defer timeCost()()
	result, tag := 0, 1
	if x < 0 {
		tag = -1
		x = tag * x
	}
	for x != 0 {
		temp := x % 10
		x = x / 10
		result = result*10 + temp
	}
	result = tag * result
	if (result < 0 && result < (-1<<31)) || (result > 0 && result > (1<<31)) {
		return 0
	}
	return result
}
