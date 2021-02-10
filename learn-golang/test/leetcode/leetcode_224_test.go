package leetcode

import (
	"fmt"
	"testing"
)

/**
 * 224. 基本计算器
 */

func Test_leetcode_224(t *testing.T) {
	fmt.Println(calculate("1-(5)"))
	fmt.Println(calculate("- (3 + (4 + 5))"))
	// fmt.Println(calculate("1 + 1"))
	// fmt.Println(calculate(" 2-1 + 2 "))
	fmt.Println(calculate("(1+(4+5+2)-3)+(6+8)"))
}

func calculate(s string) int {
	defer timeCost()()
	f := func(temp int, sign rune) int {
		if sign == '+' {
			return temp
		}
		return -temp

	}
	temp, sign, stack1, stack2, result := 0, '+', make([]int, 0), make([]rune, 0), 0
	for _, c := range s {
		if c == ' ' {
			continue
		} else if c >= '0' && c <= '9' {
			temp = temp*10 + int(c) - 48
		} else {
			result += f(temp, sign)
			temp = 0

			if c == '(' {
				stack1 = append(stack1, result)
				result = 0

				stack2 = append(stack2, sign)
				sign = '+'
			} else if c == ')' {
				len1 := len(stack1) - 1
				len2 := len(stack2) - 1

				result = stack1[len1] + f(result, stack2[len2])
				sign = '+'

				stack1 = stack1[0:len1]
				stack2 = stack2[0:len2]
			} else {
				sign = c
			}
		}
	}

	return result + f(temp, sign)
}
