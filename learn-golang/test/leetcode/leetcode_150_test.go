package leetcode

import (
	"fmt"
	"strconv"
	"testing"
)

/**
 * 150. 逆波兰表达式求值
 */

func Test_leetcode_150(t *testing.T) {
	fmt.Println(evalRPN([]string{"2", "1", "+", "3", "*"}))
	fmt.Println(evalRPN([]string{"4", "13", "5", "/", "+"}))
	fmt.Println(evalRPN([]string{"10", "6", "9", "3", "+", "-11", "*", "/", "*", "17", "+", "5", "+"}))
}

func evalRPN(tokens []string) int {
	defer timeCost()()
	f := func(w1, w2, s string) int {
		num1, _ := strconv.Atoi(w1)
		num2, _ := strconv.Atoi(w2)
		switch s {
		case "+":
			return num1 + num2
		case "-":
			return num1 - num2
		case "*":
			return num1 * num2
		case "/":
			return num1 / num2
		}
		panic("")
	}
	length := len(tokens)
	for length > 1 {
		for i := 0; i < length; i++ {
			if tokens[i] == "+" || tokens[i] == "-" || tokens[i] == "*" || tokens[i] == "/" {
				tokens = append(tokens[:i-2], append([]string{strconv.Itoa(f(tokens[i-2], tokens[i-1], tokens[i]))}, tokens[i+1:]...)...)
				break
			}
		}
		length = len(tokens)
	}
	result, _ := strconv.Atoi(tokens[0])
	return result
}
