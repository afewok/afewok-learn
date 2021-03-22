package leetcode

import (
	"fmt"
	"strconv"
	"testing"
)

/**
 * 1154. 一年中的第几天
 */

func Test_leetcode_1154(t *testing.T) {
	fmt.Println(dayOfYear("2019-01-09"))
	fmt.Println(dayOfYear("2019-02-10"))
	fmt.Println(dayOfYear("2003-03-01"))
	fmt.Println(dayOfYear("2004-03-01"))
}

func dayOfYear(date string) int {
	defer timeCost()()
	year, _ := strconv.Atoi(date[:4])
	mon, _ := strconv.Atoi(date[5:7])
	days, _ := strconv.Atoi(date[8:])
	for i := 1; i < mon; i++ {
		days += getDays(year, i)
	}
	return days
}

func getDays(year, mon int) int {
	switch mon {
	case 1:
		return 31
	case 2:
		if year%4 == 0 && (year%100 != 0 || year%400 == 0) {
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
	}
	return 0
}
