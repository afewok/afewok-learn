package leetcode

import (
	"fmt"
	"testing"
)

/**
 * 69. x 的平方根
 */

func Test_leetcode_069(t *testing.T) {
	fmt.Println(mySqrt(1))
	fmt.Println(mySqrt(4))
	fmt.Println(mySqrt(8))
}

func mySqrt(x int) int {
	defer timeCost()()
	left, right, mid, temp := 0, x, 0, 0
	for left < right {
		mid = left + (right-left)/2
		temp = mid * mid
		if temp == x {
			return mid
		} else if temp > x {
			right = mid
		} else {
			temp = (mid + 1) * (mid + 1)
			if temp == x {
				return mid + 1
			}
			if temp > x {
				return mid
			}
			left = mid
		}
	}
	return 0
}
