package leetcode

import (
	"fmt"
	"testing"
)

/**
 * 868. 二进制间距
 */

func Test_leetcode_868(t *testing.T) {
	fmt.Println(binaryGap(22))
	fmt.Println(binaryGap(5))
	fmt.Println(binaryGap(6))
	fmt.Println(binaryGap(8))
}

func binaryGap(n int) int {
	defer timeCost()()
	max, count := 0, -1
	for n > 0 {
		if count > -1 {
			count++
		}
		if n&1 == 1 {
			if max < count {
				max = count
			}
			count = 0
		}
		n = n >> 1
	}
	return max
}
