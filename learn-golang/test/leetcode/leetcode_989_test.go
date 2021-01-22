package leetcode

import (
	"fmt"
	"testing"
)

/**
 * 989. 数组形式的整数加法
 */

func Test_leetcode_989(t *testing.T) {
	fmt.Println(addToArrayForm([]int{1, 2, 0, 0}, 34))
}

func addToArrayForm(A []int, K int) []int {
	defer timeCost()()
	plusOne, count := 0, len(A)-1
	for plusOne > 0 || K > 0 {
		temp := 0
		if K > 0 {
			temp = K % 10
			K = K / 10
		}

		if count < 0 {
			A = append([]int{plusOne + temp}, A...)
			plusOne = A[0] / 10
			A[0] = A[0] % 10
		} else {
			A[count] = A[count] + plusOne + temp
			plusOne = A[count] / 10
			A[count] = A[count] % 10
			count--
		}
	}
	return A
}
