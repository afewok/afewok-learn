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
	temp, result := 0, make([]bool, len(A))
	for i, num := range A {
		temp = (temp*2 + num) % 10
		if temp == 0 || temp == 5 {
			result[i] = true
		}
	}
	return result
}

//1，2，4，8，16
//110、1100、11000
//6、12，20
