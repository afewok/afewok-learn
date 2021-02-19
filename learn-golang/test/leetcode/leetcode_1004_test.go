package leetcode

import (
	"fmt"
	"testing"
)

/**
 * 1004. 最大连续1的个数 III
 */

func Test_leetcode_1004(t *testing.T) {
	fmt.Println(longestOnes([]int{0, 0, 0}, 0))
	fmt.Println(longestOnes([]int{1, 1, 1, 0, 0, 0, 1, 1, 1, 1, 0, 0, 0, 0, 0}, 2))
	fmt.Println(longestOnes([]int{0, 0, 1, 1, 0, 0, 1, 1, 1, 0, 1, 1, 0, 0, 0, 1, 1, 1, 1}, 3))
}

func longestOnes(A []int, K int) int {
	defer timeCost()()
	count, sub, max := 0, 0, 0
	for i, v := range A {
		if v == 0 {
			if K == 0 {
				sub = i + 1
			} else if count < K {
				count++
			} else {
				for sub < i {
					if A[sub] == 0 {
						sub++
						break
					}
					sub++
				}
			}
		}
		if max < i-sub+1 {
			max = i - sub + 1
		}
	}
	return max
}

func longestOnes1(A []int, K int) int {
	defer timeCost()()
	length, left, zeroCount := len(A), 0, 0
	for i := 0; i < length; i++ {
		zeroCount += 1 - A[i]
		if zeroCount > K {
			zeroCount -= 1 - A[left]
			left++
		}
	}
	return length - left
}

func longestOnes2(A []int, K int) int {
	left, zeroCnt := 0, 0
	for _, num := range A {
		if num == 0 {
			zeroCnt++
		}
		if zeroCnt > K {
			if A[left] == 0 {
				zeroCnt--
			}
			left++
		}
	}
	return len(A) - left
}
