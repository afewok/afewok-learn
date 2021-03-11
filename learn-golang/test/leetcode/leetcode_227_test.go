package leetcode

import (
	"fmt"
	"strconv"
	"testing"
)

/**
 * 227. 基本计算器 II
 */

func Test_leetcode_227(t *testing.T) {
	// fmt.Println(calculate227("3+2*2"))
	// fmt.Println(calculate227(" 3/2 "))
	// fmt.Println(calculate227(" 3+5 / 2 "))
	// fmt.Println(calculate227(""))
	// fmt.Println(calculate227("0"))
	fmt.Println(calculate227("1+1+1"))
}

func calculate227(s string) int {
	defer timeCost()()
	compute := func(a, b int, symbol byte) int {
		switch symbol {
		case '+':
			return a + b
		case '-':
			return a - b
		case '*':
			return a * b
		case '/':
			return a / b
		}
		panic("panic")
	}
	length, temp, stack1, stack2 := len(s), make([]byte, 0), make([]int, 0), make([]byte, 0)
	for i := 0; i < length; i++ {
		if s[i] >= '0' && s[i] <= '9' {
			temp = append(temp, s[i])
		} else if s[i] == '+' || s[i] == '-' || s[i] == '*' || s[i] == '/' {
			num, _ := strconv.Atoi(string(temp))
			stack1 = append(stack1, num)
			if len(stack2) == 1 {
				if (stack2[0] == '*' || stack2[0] == '/') || (s[i] == '+' || s[i] == '-') {
					stack1[0] = compute(stack1[0], stack1[1], stack2[0])
					stack1 = stack1[:1]
					stack2 = stack2[:0]
				}
			} else if len(stack2) == 2 {
				stack1[1] = compute(stack1[1], stack1[2], stack2[1])
				stack1 = stack1[:2]
				stack2 = stack2[:1]
				if (stack2[0] == '*' || stack2[0] == '/') || (s[i] == '+' || s[i] == '-') {
					stack1[0] = compute(stack1[0], stack1[1], stack2[0])
					stack1 = stack1[:1]
					stack2 = stack2[:0]
				}
			}

			temp = temp[:0]
			stack2 = append(stack2, s[i])
		}
	}
	if len(temp) > 0 {
		num, _ := strconv.Atoi(string(temp))
		if len(stack2) == 0 {
			return num
		} else if len(stack2) == 1 {
			return compute(stack1[0], num, stack2[0])
		}
		return compute(stack1[0], compute(stack1[1], num, stack2[1]), stack2[0])
	}
	return 0
}
