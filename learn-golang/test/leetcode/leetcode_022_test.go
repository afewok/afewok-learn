package leetcode

import (
	"fmt"
	"testing"
)

/**
 * 22. 括号生成
 */
func Test_leetcode_022(t *testing.T) {
	fmt.Println(generateParenthesis(1))
	fmt.Println(generateParenthesis(2))
	fmt.Println(generateParenthesis(3))
	fmt.Println(generateParenthesis(8))
}

func generateParenthesis(n int) []string {
	defer timeCost()()
	f1 := func(length int) []string {
		result := make([]string, 0)
		result = append(result, "(", ")")
		for i := 1; i < length; i++ {
			count, temp := len(result), make([]string, 0)
			for j := 0; j < count; j++ {
				temp = append(temp, result[j]+"(")
				temp = append(temp, result[j]+")")
			}
			result = temp
		}
		return result
	}
	f2 := func(str string) bool {
		count := 0
		for _, v := range str {
			if v == '(' {
				count++
			} else {
				count--
			}
			if count < 0 {
				return false
			}
		}
		if count != 0 {
			return false
		}
		return true
	}
	strs, mp, result := f1(2*n), make(map[string]bool), make([]string, 0)
	for _, str := range strs {
		if f2(str) {
			if _, OK := mp[str]; !OK {
				mp[str] = true
				result = append(result, str)
			}
		}
	}
	return result
}
