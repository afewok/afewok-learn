package leetcode

import (
	"fmt"
	"testing"
)

/**
 * 1185. 一周中的第几天
 */

func Test_leetcode_1185(t *testing.T) {
	fmt.Println(dayOfTheWeek(31, 8, 2019))
	fmt.Println(dayOfTheWeek(18, 7, 1999))
	fmt.Println(dayOfTheWeek(15, 8, 1993))
}

func dayOfTheWeek(day int, month int, year int) string {
	defer timeCost()()
	return ""
}
