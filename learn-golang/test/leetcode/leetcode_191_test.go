package leetcode

import (
	"fmt"
	"testing"
)

/**
 * 191. 位1的个数
 */
func Test_leetcode_191(t *testing.T) {
	fmt.Println(hammingWeight(43261596))
	fmt.Println(hammingWeight(4294967293))
}

func hammingWeight(num uint32) int {
	defer timeCost()()
	result := 0
	for num > 0 {
		if num%2 == 1 {
			result++
		}
		num = num >> 1
	}
	return result
}
