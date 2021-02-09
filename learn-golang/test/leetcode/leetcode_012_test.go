package leetcode

import (
	"fmt"
	"testing"
)

/**
 * 12. 整数转罗马数字
 */

func Test_leetcode_012(t *testing.T) {
	fmt.Println(intToRoman(3))
	fmt.Println(intToRoman(4))
	fmt.Println(intToRoman(9))
	fmt.Println(intToRoman(58))   //LVIII
	fmt.Println(intToRoman(1994)) //MCMXCIV
}

func intToRoman(num int) string {
	defer timeCost()()
	f := func(a1, a5, a10 string, count int) string {
		result := ""
		if (count > 0 && count < 4) || (count > 4 && count < 9) {
			if count >= 5 {
				count -= 5
				result += a5
			}
			for i := 0; i < count; i++ {
				result += a1
			}
		} else if count == 4 {
			result += a1 + a5
		} else if count == 9 {
			result += a1 + a10
		}
		return result
	}

	result := ""
	count := num / 1000
	num = num % 1000
	for i := 0; i < count; i++ {
		result += "M"
	}
	count = num / 100
	num = num % 100
	result += f("C", "D", "M", count)
	count = num / 10
	num = num % 10
	result += f("X", "L", "C", count)
	result += f("I", "V", "X", num)
	return result
}
