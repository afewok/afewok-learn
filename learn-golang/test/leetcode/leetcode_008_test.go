package leetcode

import (
	"fmt"
	"testing"
)

/**
 * 8. 字符串转换整数 (atoi)
 */
func Test_leetcode_008(t *testing.T) {
	fmt.Println(myAtoi("2147483648"))
	fmt.Println(myAtoi("1095502006p8"))
	fmt.Println(myAtoi("2147483646"))
	fmt.Println(myAtoi("2147483647"))
	fmt.Println(myAtoi("  0000000000012345678"))
	fmt.Println(myAtoi("9223372036854775808"))
	fmt.Println(myAtoi("42"))
	fmt.Println(myAtoi("   -42"))
	fmt.Println(myAtoi("4193 with words"))
	fmt.Println(myAtoi("words and 987"))
	fmt.Println(myAtoi("-91283472332"))
}

func myAtoi(s string) int {
	defer timeCost()()
	f := func(s string) []rune {
		isStart, result := false, make([]rune, 0)
		for _, c := range s {
			if isStart {
				if c < '0' || c > '9' {
					break
				} else if c == '0' && len(result) == 1 {
					continue
				}
				result = append(result, c)
			} else if c == ' ' {
				continue
			} else if c == '+' || c == '-' {
				isStart = !isStart
				result = append(result, c)
			} else if c < '0' || c > '9' {
				result = append(result, '0')
				break
			} else {
				isStart = !isStart
				if c == '0' {
					result = append(result, '+')
				} else {
					result = append(result, '+', c)
				}
			}
		}
		return result
	}
	f1 := func(temp []rune) int {
		result, length := 0, len(temp)
		for i := 1; i < length; i++ {
			result = result*10 + int(temp[i]) - 48
		}
		return result
	}
	f2 := func(sub, symbol, def int, str string, temp []rune) int {
		for i := 0; i < 10; i++ {
			if int(str[i]) > int(temp[sub+i]) {
				break
			} else if int(str[i]) < int(temp[sub+i]) {
				return symbol * def
			}
		}
		return symbol * f1(temp)
	}
	temp := f(s)
	length := len(temp)
	if length < 2 {
		return 0
	} else if length < 11 {
		if temp[0] == '-' {
			return f1(temp) * -1
		}
		return f1(temp)
	} else if length > 11 {
		if temp[0] == '-' {
			return -2147483648
		}
		return 2147483647
	} else if temp[0] == '-' {
		return f2(length-10, -1, 2147483648, "2147483648", temp)
	}
	return f2(length-10, 1, 2147483647, "2147483647", temp)
}
