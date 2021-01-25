package leetcode

import (
	"fmt"
	"testing"
)

/**
 * 476. 数字的补数
 */

func Test_leetcode_476(t *testing.T) {
	fmt.Println(findComplement(5))
	fmt.Println(findComplement(1))
}

func findComplement(num int) int {
	defer timeCost()()
	max := 1
	for max < num {
		max = max<<1 + 1
	}
	return max ^ num
}
