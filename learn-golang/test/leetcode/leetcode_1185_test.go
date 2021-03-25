package leetcode

import (
	"fmt"
	"testing"
)

/**
 * 1185. 一周中的第几天
 */

func Test_leetcode_1185(t *testing.T) {
	fmt.Println(dayOfTheWeek(1, 1, 1970))
	fmt.Println(dayOfTheWeek(31, 8, 2019))
	fmt.Println(dayOfTheWeek(18, 7, 1999))
	fmt.Println(dayOfTheWeek(15, 8, 1993))
}

func dayOfTheWeek(day int, month int, year int) string {
	defer timeCost()()
	arr := []string{"Monday", "Tuesday", "Wednesday", "Thursday", "Friday", "Saturday", "Sunday"}
	result := 0
	for Y := 1970; Y < year; Y++ {
		if Y%4 == 0 && (Y%100 != 0 || Y%400 == 0) {
			result += 366
		} else {
			result += 365
		}
	}
	for i := 0; i < month; i++ {
		result += getDays1185(year, i)
	}
	return arr[(result+day+2)%7]
}

func getDays1185(Y int, M int) int {
	switch M {
	case 0:
		return 0
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
