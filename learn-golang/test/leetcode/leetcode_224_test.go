package leetcode

import (
	"fmt"
	"strconv"
	"testing"
)

/**
 * 224. 基本计算器
 */

func Test_leetcode_224(t *testing.T) {
	fmt.Println(calculate("1 + 1"))
	fmt.Println(calculate(" 2-1 + 2 "))
	fmt.Println(calculate("(1+(4+5+2)-3)+(6+8)"))
}

func calculate(s string) int {
	defer timeCost()()
	temp, stack, result := "", make([]int, 0), 0
	for _, c := range s {
		if c == ' ' {
			continue
		} else if c >= '0' && c <= '9' {
			temp = string(append([]rune(temp), c))
		} else if c == '(' {
			stack = append(stack, result)
			result = 0
		}

		switch c {
		case ' ':
		case '(':

		case ')':
			length := len(stack) - 1
			result += stack[length]
			stack = stack[0:length]
		case '+':
			n, _ := strconv.Atoi(temp)
			result += n
		case '-':
			n, _ := strconv.Atoi(temp)
			result -= n
		default:

		}
	}
	return result
}
