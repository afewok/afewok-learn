package leetcode

import (
	"fmt"
	"testing"
)

/**
 * 441. 排列硬币
 */

func Test_leetcode_441(t *testing.T) {
	fmt.Println(arrangeCoins(1))
	fmt.Println(arrangeCoins(5))
	fmt.Println(arrangeCoins(8))
}

func arrangeCoins(n int) int {
	defer timeCost()()
	if n == 0 {
		return 0
	}
	count := 1
	for n >= count {
		n -= count
		count++
	}
	return count - 1
}
