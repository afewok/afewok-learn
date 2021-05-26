package leetcode

import (
	"fmt"
	"testing"
)

/**
 * 1190. 反转每对括号间的子串
 */

func Test_leetcode_1190(t *testing.T) {
	// fmt.Println(reverseParentheses("(abcd)"))
	// fmt.Println(reverseParentheses("(u(love)i)"))
	// fmt.Println(reverseParentheses("(ed(et(oc))el)"))
	fmt.Println(reverseParentheses("a(bcdefghijkl(mno)p)q"))
	// fmt.Println(reverseParentheses("yfgnxf"))
	// fmt.Println(reverseParentheses("ta()usw((((a))))"))
}

func reverseParentheses(s string) string {
	defer timeCost()()
	arrs, lefts := make([]rune, 0), make([]int, 0)
	for _, v := range s {
		if v == '(' {
			lefts = append(lefts, len(arrs))
		} else if v == ')' {
			left, right := lefts[len(lefts)-1], len(arrs)-1
			for left < right {
				arrs[left], arrs[right] = arrs[right], arrs[left]
				left++
				right--
			}

			lefts = lefts[:len(lefts)-1]
		} else {
			arrs = append(arrs, v)
		}
	}

	return string(arrs)
}
