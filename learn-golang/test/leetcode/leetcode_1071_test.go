package leetcode

import (
	"fmt"
	"testing"
)

/**
 * 1071. 字符串的最大公因子
 */

func Test_leetcode_1071(t *testing.T) {
	fmt.Println(gcdOfStrings("ABCABC", "ABC"))
	fmt.Println(gcdOfStrings("ABABAB", "ABAB"))
	fmt.Println(gcdOfStrings("LEET", "CODE"))
}

func gcdOfStrings(str1 string, str2 string) string {
	defer timeCost()()
	result, len1, len2 := "", len(str1), len(str2)
	if len1 == len2 {
		if str1 == str2 {
			return str1
		}
		return ""
	} else if len1 > len2 {
		result = str1[len2:len1]
	} else if len1 < len2 {
		result = str2[len1:len2]
	}

	len3 := len(result)
	for i := len3; i > 0; i-- {
		if len1%i != 0 || len2%i != 0 || len3%i != 0 {
			continue
		}
		r1, r2 := len1/i, len2/i
		temp1, temp2, temp3 := "", "", result[0:i]
		for i := 0; i < r1; i++ {
			temp1 += temp3
		}
		for i := 0; i < r2; i++ {
			temp2 += temp3
		}
		if str1 == temp1 && str2 == temp2 {
			return temp3
		}
	}
	return ""
}
