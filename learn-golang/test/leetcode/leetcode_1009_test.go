package leetcode

import (
	"fmt"
	"testing"
)

/**
 * 1009. 十进制整数的反码
 */

func Test_leetcode_1009(t *testing.T) {
	// fmt.Println(bitwiseComplement(5))
	// fmt.Println(bitwiseComplement(7))
	fmt.Println(bitwiseComplement(10))
}

func bitwiseComplement(N int) int {
	defer timeCost()()
	result := 1
	for result < N {
		result = (result << 1) + 1
	}
	return result ^ N
}
