package leetcode

import (
	"fmt"
	"math"
	"testing"
)

/**
 * 367. 有效的完全平方数
 */

func Test_leetcode_367(t *testing.T) {
	fmt.Println(isPerfectSquare(16))
	fmt.Println(isPerfectSquare(14))
}

func isPerfectSquare(num int) bool {
	defer timeCost()()
	max := num/2 + 1
	for i := 1; i <= max; i++ {
		temp := i * i
		if temp == num {
			return true
		} else if temp > num {
			return false
		}
	}
	return false
}
func isPerfectSquare1(num int) bool {
	defer timeCost()()
	temp := int(math.Sqrt(float64(num)))
	if temp*temp == num {
		return true
	}
	return false
}
