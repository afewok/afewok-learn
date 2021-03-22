package leetcode

import (
	"fmt"
	"strconv"
	"testing"
)

/**
 * 1360. 日期之间隔几天
 */

func Test_leetcode_1360(t *testing.T) {
	fmt.Println(daysBetweenDates("2019-06-29", "2019-06-30"))
	fmt.Println(daysBetweenDates("2020-01-15", "2019-12-31"))
}

func daysBetweenDates(date1 string, date2 string) int {
	defer timeCost()()
	year1, _ := strconv.Atoi(date1[:4])
	mon1, _ := strconv.Atoi(date1[5:7])
	days1, _ := strconv.Atoi(date1[8:])

	year2, _ := strconv.Atoi(date2[:4])
	mon2, _ := strconv.Atoi(date2[5:7])
	days2, _ := strconv.Atoi(date2[8:])
	if year1*10000+mon1*100+days1 > year2*10000+mon2*100+days2 {
		year1, mon1, days1, year2, mon2, days2 = year2, mon2, days2, year1, mon1, days1
	}
	result := 0
	if year1 == year2 {
		if mon1 == mon2 {
			return days2 - days1
		} else {
			result += getDays1360(year1, mon1) - days1
			result += days2
			for i := mon1 + 1; i < mon2; i++ {
				result += getDays1360(year1, i)
			}
		}
	} else {
		result += getDays1360(year1, mon1) - days1
		for i := mon1 + 1; i <= 12; i++ {
			result += getDays1360(year1, i)
		}
		for i := year1 + 1; i < year2; i++ {
			result += getYear1360(i)
		}

		result += days2
		for i := 1; i < mon2; i++ {
			result += getDays1360(year2, i)
		}
	}
	return result
}
func getYear1360(year int) int {
	if year%4 == 0 && (year%100 != 0 || year%400 == 0) {
		return 366
	}
	return 365
}

func getDays1360(year, mon int) int {
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
	case 12:
		return 31
	}
	return 0
}
