package leetcode

import (
	"fmt"
	"testing"
)

/**
 * 5649. 解码异或后的数组
 */

func Test_leetcode_5649(t *testing.T) {
	fmt.Println(decode([]int{1, 2, 3}, 1))
	fmt.Println(decode([]int{6, 2, 7, 3}, 4))
}

func decode(encoded []int, first int) []int {
	defer timeCost()()
	temp, length := encoded[0], len(encoded)
	encoded[0] = first
	for i := 1; i < length; i++ {
		temp, encoded[i] = encoded[i], encoded[i-1]^temp
	}
	return append(encoded, encoded[length-1]^temp)
}
