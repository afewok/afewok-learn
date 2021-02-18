package leetcode

import (
	"fmt"
	"testing"
)

/**
 * 995. K 连续位的最小翻转次数
 * 思路：差分数组、滑动窗口
 */
func Test_leetcode_995(t *testing.T) {
	// fmt.Println(minKBitFlips([]int{0, 1, 0}, 1))
	// fmt.Println(minKBitFlips([]int{1, 1, 0}, 2))
	fmt.Println(minKBitFlips([]int{0, 0, 0, 1, 0, 1, 1, 0}, 3))
	// fmt.Println(minKBitFlips([]int{0, 1, 0, 1, 1, 1, 0, 1}, 3))
}

func minKBitFlips1(A []int, K int) int {
	defer timeCost()()
	length := len(A)
	temp, result, num := make([]int, length+1), 0, 0
	for i := 0; i < length; i++ {
		num ^= temp[i]
		if A[i] == num {
			if i+K > length {
				return -1
			}
			result++
			num ^= 1
			temp[i+K] = 1
		}
	}
	return result
}

func minKBitFlips(A []int, K int) int {
	defer timeCost()()
	length, num, result := len(A), 0, 0
	for i := 0; i < length; i++ {
		if i >= K && A[i-K] < 0 {
			num ^= 1
			A[i-K] += 2
		}
		if A[i] == num {
			if i+K > length {
				return -1
			}
			result++
			num ^= 1
			A[i] -= 2
		}
	}
	return result
}
