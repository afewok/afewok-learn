package leetcode

import (
	"fmt"
	"testing"
)

/**
 * 1018. 可被 5 整除的二进制前缀
 */

func Test_leetcode_1018(t *testing.T) {
	// fmt.Println(prefixesDivBy5([]int{0, 1, 1}))
	// fmt.Println(prefixesDivBy5([]int{1, 1, 1}))
	// fmt.Println(prefixesDivBy5([]int{0, 1, 1, 1, 1, 1})) //32，16、8、4、2、1
	// fmt.Println(prefixesDivBy5([]int{1, 1, 1, 0, 1}))
	fmt.Println(prefixesDivBy5([]int{1, 1, 0, 1, 1, 1, 1, 1, 1, 1, 1, 0, 1, 1, 1, 0, 0, 1, 0}))
}

func prefixesDivBy5(A []int) []bool {
	defer timeCost()()
	temp, result, arr := 0, make([]bool, len(A)), []int{2, 4, 8, 6}
	for i, num := range A {
		temp = num
		for j := 1; j <= i; j++ {
			if A[i-j] == 1 {
				temp += arr[(j-1)%4]
			}
		}
		if temp%5 == 0 {
			result[i] = true
		}
	}
	return result
}
