package leetcode

import (
	"fmt"
	"testing"
)

/**
 * 709. 转换成小写字母
 *
 */
func Test_leetcode_709(t *testing.T) {
	fmt.Println(toLowerCase("Hello"))
	fmt.Println(toLowerCase("here"))
	fmt.Println(toLowerCase("LOVELY"))
}

func toLowerCase(str string) string {
	defer timeCost()()
	result := make([]byte, len(str))
	for i, v1 := range str {
		if v1 >= 65 && v1 <= 90 {
			result[i] = byte(v1 + 32)
		} else {
			result[i] = byte(v1)
		}
	}
	return string(result)
}
