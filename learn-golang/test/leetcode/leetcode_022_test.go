package leetcode

import (
	"fmt"
	"testing"
)

/**
 * 22. 括号生成
 */
func Test_leetcode_022(t *testing.T) {
	fmt.Println(generateParenthesis(3))
}

func generateParenthesis(n int) []string {
	defer timeCost()()
	return []string{}
}
