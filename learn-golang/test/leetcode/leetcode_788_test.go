package leetcode

import (
	"fmt"
	"testing"
)

/**
 * 788. 旋转数字
 *
 */
func Test_leetcode_788(t *testing.T) {
	fmt.Println(rotatedDigits(10))
}

func rotatedDigits(N int) int {
	defer timeCost()()
	temp1, temp2, isRotated, count, result, mp := 0, 0, true, 0, 0, map[int]int{0: 0, 1: 1, 2: 5, 5: 2, 6: 9, 8: 8, 9: 6}
	for i := 0; i <= N; i++ {
		temp1 = i
		temp2 = 0
		count = 0
		isRotated = true
		for temp1 > 0 {
			temp := temp1 % 10
			if v, OK := mp[temp]; OK {
				for j := 0; j < count; j++ {
					v *= 10
				}
				temp2 += v
				temp1 /= 10
				count++
			} else {
				isRotated = false
				break
			}

		}
		if isRotated && temp2 != i {
			result++
		}
	}
	return result
}

//0，1，8
//2，5，6，9
//3，4，7

//25
