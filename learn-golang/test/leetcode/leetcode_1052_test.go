package leetcode

import (
	"fmt"
	"testing"
)

/**
 * 1052. 爱生气的书店老板
 */

func Test_leetcode_1052(t *testing.T) {
	fmt.Println(maxSatisfied([]int{1, 0, 1, 2, 1, 1, 7, 5}, []int{0, 1, 0, 1, 0, 1, 0, 1}, 3))
}

func maxSatisfied(customers []int, grumpy []int, X int) int {
	defer timeCost()()
	length, result, temp := len(customers), 0, 0
	for i := 0; i < X; i++ {
		if i == length {
			return result + temp
		}
		if grumpy[i] == 0 {
			result += customers[i]
		} else {
			temp += customers[i]
		}
	}
	max := temp
	for i := X; i < length; i++ {
		if grumpy[i] == 0 {
			result += customers[i]
		} else {
			temp += customers[i]
		}
		if grumpy[i-X] != 0 {
			temp -= customers[i-X]
		}
		if max < temp {
			max = temp
		}
	}
	return result + max
}
