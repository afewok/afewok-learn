package leetcode

import (
	"fmt"
	"testing"
)

/**
 * 1118. 一月有多少天
 */

func Test_leetcode_1118(t *testing.T) {
	fmt.Println(numberOfDays(1992, 7))
	fmt.Println(numberOfDays(1900, 2))
	fmt.Println(numberOfDays(2000, 2))
}

func numberOfDays(Y int, M int) int {
	defer timeCost()()
	switch M {
	case 1:
		return 31
	case 2:
		if Y%4 == 0 && (Y%100 != 0 || Y%400 == 0) {
			return 29
		}
		return 28
	case 3:
		return 31
	case 4:
		return 30
	case 5:
		return 31
	case 6:
		return 30
	case 7:
		return 31
	case 8:
		return 31
	case 9:
		return 30
	case 10:
		return 31
	case 11:
		return 30
	case 12:
		return 31
	}
	panic("panic panic panic")
}
