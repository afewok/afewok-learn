package leetcode

import (
	"fmt"
	"testing"
)

/**
 * 338. 比特位计数
 */

func Test_leetcode_338(t *testing.T) {
	fmt.Println(countBits(2))
	fmt.Println(countBits(5))
}

//方法一：直接计算
func countBits1(num int) []int {
	defer timeCost()()
	result := make([]int, num+1)
	result[0] = 0
	for i := 1; i <= num; i++ {
		temp, n := 0, i
		for n > 0 {
			if n%2 == 1 {
				temp++
			}
			n = n / 2
		}
		result[i] = temp
	}
	return result
}

//方法二：动态规划——最高有效位
func countBits(num int) []int {
	bits := make([]int, num+1)
	highBit := 0
	for i := 1; i <= num; i++ {
		if i&(i-1) == 0 {
			highBit = i
		}
		bits[i] = bits[i-highBit] + 1
	}
	return bits
}

//方法三：动态规划——最低有效位
func countBits3(num int) []int {
	bits := make([]int, num+1)
	for i := 1; i <= num; i++ {
		bits[i] = bits[i>>1] + i&1
	}
	return bits
}

//方法四：动态规划——最低设置位
func countBits4(num int) []int {
	bits := make([]int, num+1)
	for i := 1; i <= num; i++ {
		bits[i] = bits[i&(i-1)] + 1
	}
	return bits
}
