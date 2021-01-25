package leetcode

import (
	"fmt"
	"testing"
)

/**
 * 5661. 替换隐藏数字得到的最晚时间
 */

func Test_leetcode_5661(t *testing.T) {
	fmt.Println(maximumTime("2?:?0"))
	fmt.Println(maximumTime("0?:3?"))
	fmt.Println(maximumTime("1?:22"))
}

func maximumTime(time string) string {
	defer timeCost()()
	result := make([]byte, 5)
	if time[0] == '?' {
		if time[1] == '?' || time[1] < '4' {
			result[0] = '2'
		} else {
			result[0] = '1'
		}
	} else {
		result[0] = time[0]
	}
	if time[1] == '?' {
		if result[0] == '0' || result[0] == '1' {
			result[1] = '9'
		} else {
			result[1] = '3'
		}
	} else {
		result[1] = time[1]
	}
	result[2] = ':'
	if time[3] == '?' {
		result[3] = '5'
	} else {
		result[3] = time[3]
	}
	if time[4] == '?' {
		result[4] = '9'
	} else {
		result[4] = time[4]
	}
	return string(result)
}
