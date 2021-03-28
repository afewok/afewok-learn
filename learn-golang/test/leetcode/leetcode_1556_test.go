package leetcode

import (
	"fmt"
	"strconv"
	"testing"
)

/**
 * 1556. 千位分隔数
 */

func Test_leetcode_1556(t *testing.T) {
	fmt.Println(thousandSeparator(987))
	fmt.Println(thousandSeparator(1234))
	fmt.Println(thousandSeparator(123456789))
}

func thousandSeparator(n int) string {
	defer timeCost()()
	if n == 0 {
		return "0"
	}
	result := ""
	for n > 0 {
		if len(result) != 0 {
			result = "." + result
		}
		temp := n % 1000
		n = n / 1000
		if n > 0 && temp < 10 {
			result = "00" + strconv.Itoa(temp) + result
		} else if n > 0 && temp < 100 {
			result = "0" + strconv.Itoa(temp) + result
		} else {
			result = strconv.Itoa(temp) + result
		}
	}
	return result
}
