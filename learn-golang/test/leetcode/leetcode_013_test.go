package leetcode

import (
	"fmt"
	"testing"
)

/**
 * 13. 罗马数字转整数
 */
func Test_leetcode_013(t *testing.T) {
	fmt.Println(romanToInt("IXCM"))
	fmt.Println(romanToInt("III"))
	fmt.Println(romanToInt("IV"))
	fmt.Println(romanToInt("IX"))
	fmt.Println(romanToInt("LVIII"))
	fmt.Println(romanToInt("MCMXCIV"))
	fmt.Println(romanToInt("MMMM"))
}

func romanToInt(s string) int {
	defer timeCost()()
	mp1 := map[byte]int{'I': 1, 'V': 5, 'X': 10, 'L': 50, 'C': 100, 'D': 500, 'M': 1000}
	mp2 := map[byte]int{'I': 1, 'V': 2, 'X': 3, 'L': 4, 'C': 5, 'D': 6, 'M': 7}
	result, length := 0, len(s)

	if v, OK := mp1[s[length-1]]; OK {
		result += v
	}
	for i := length - 2; i >= 0; i-- {
		index1, _ := mp2[s[i+1]]
		index2, _ := mp2[s[i]]
		v2, _ := mp1[s[i]]
		if index1 > index2 {
			result -= v2
		} else {
			result += v2
		}
	}
	return result
}
