package leetcode

import (
	"fmt"
	"testing"
)

/**
 * 22. 括号生成
 * 思路：暴力，深度优先遍历（递归+回溯+剪枝）、广度优先遍历（动态规划（杨辉三角第n行））
 */
func Test_leetcode_022(t *testing.T) {
	// fmt.Println(generateParenthesis(1))
	fmt.Println(generateParenthesis(2))
	// fmt.Println(generateParenthesis(3))
	// fmt.Println(generateParenthesis(8))
}

func generateParenthesis1(n int) []string {
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

func generateParenthesis2(n int) []string {
	result := make([]string, 0)
	getParenthesis("", n, n, &result)
	return result
}

func getParenthesis(str string, left, right int, result *[]string) {
	if left == 0 && right == 0 {
		*result = append(*result, str)
		return
	}

	if left == right {
		getParenthesis(str+"(", left-1, right, result)
	} else if left < right {
		if left > 0 {
			getParenthesis(str+"(", left-1, right, result)
		}
		getParenthesis(str+")", left, right-1, result)
	}
}

//广度优先（杨辉三角）
func generateParenthesis(n int) []string {
	result := make([][]string, n+1)
	result[0] = []string{""}
	for i := 1; i <= n; i++ {
		result[i] = make([]string, 0)
		for j := 0; j < i; j++ {
			for _, intside := range result[j] {
				for _, outside := range result[i-j-1] {
					result[i] = append(result[i], "("+intside+")"+outside)
				}
			}
		}
	}
	return result[n]
}
