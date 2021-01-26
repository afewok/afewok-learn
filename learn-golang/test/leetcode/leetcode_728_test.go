package leetcode

import (
	"fmt"
	"testing"
)

/**
 * 728. 自除数
 *
 * 思路：
 */
func Test_leetcode_728(t *testing.T) {
	fmt.Println(selfDividingNumbers(1, 22))
}

func selfDividingNumbers(left int, right int) []int {
	defer timeCost()()
	result, isAdd := make([]int, 0), true
	for i := left; i <= right; i++ {
		if i == 0 {
			continue
		}
		isAdd = true
		num := i
		if num < 0 {
			num = -1 * num
		}
		temp := num
		for temp > 0 {
			a := temp % 10
			if a == 0 || num%a != 0 {
				isAdd = false
				break
			}
			temp = temp / 10
		}
		if isAdd {
			result = append(result, i)
		}
	}
	return result
}
