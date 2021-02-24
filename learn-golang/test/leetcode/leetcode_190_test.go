package leetcode

import (
	"fmt"
	"testing"
)

/**
 * 190. 颠倒二进制位
 */
func Test_leetcode_190(t *testing.T) {
	fmt.Println(reverseBits(43261596))
	fmt.Println(reverseBits(4294967293))
}

func reverseBits(num uint32) uint32 {
	defer timeCost()()
	var result uint32 = 0
	for i := 0; i < 32; i++ {
		if num > 0 {
			if num%2 == 1 {
				result = result<<1 + 1
			} else {
				result = result << 1
			}
			num = num / 2
		} else {
			result = result << 1
		}
	}
	return result
}
