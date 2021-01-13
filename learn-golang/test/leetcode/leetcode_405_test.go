package leetcode

import (
	"fmt"
	"testing"
)

/**
 * 405. 数字转换为十六进制数
 */

func Test_leetcode_405(t *testing.T) {
	fmt.Println(toHex(26))
	fmt.Println(toHex(1))
	fmt.Println(toHex(3))
}

func toHex(num int) string {
	defer timeCost()()
	if num == 0 {
		return "0"
	}
	result, mp := "", map[int]string{1: "1", 2: "2", 3: "3", 4: "4", 5: "5", 6: "6", 7: "7", 8: "8", 9: "9", 10: "a", 11: "b", 12: "c", 13: "d", 14: "e", 15: "f"}
	for num != 0 {
		for k, v := range mp {
			if num|k == num && num&k == k {
				result = v + result
				break
			}
		}
		num = num >> 4
	}
	return result
}

//8,4,2,1
