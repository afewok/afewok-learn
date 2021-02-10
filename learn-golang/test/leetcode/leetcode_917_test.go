package leetcode

import (
	"fmt"
	"testing"
)

/**
 * 917. 仅仅反转字母
 */
func Test_leetcode_703(t *testing.T) {
	fmt.Println(reverseOnlyLetters("ab-cd"))
	fmt.Println(reverseOnlyLetters("a-bC-dEf-ghIj"))
	fmt.Println(reverseOnlyLetters("Test1ng-Leet=code-Q!"))
}

func reverseOnlyLetters(S string) string {
	defer timeCost()()
	temp, count, result := make([]rune, 0), 0, make([]rune, 0)
	for _, s := range S {
		if (s >= 'A' && s <= 'Z') || (s >= 'a' && s <= 'z') {
			temp = append(temp, s)
		}
	}
	count = len(temp) - 1
	for _, s := range S {
		if (s >= 'A' && s <= 'Z') || (s >= 'a' && s <= 'z') {
			result = append(result, temp[count])
			count--
		} else {
			result = append(result, s)
		}
	}
	return string(result)
}
