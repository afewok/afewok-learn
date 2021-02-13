package leetcode

import (
	"fmt"
	"testing"
)

/**
 * 17. 电话号码的字母组合
 */

func Test_leetcode_017(t *testing.T) {
	fmt.Println(letterCombinations("23"))
	fmt.Println(letterCombinations(""))
	fmt.Println(letterCombinations("2"))
	fmt.Println(letterCombinations("22"))
}

func letterCombinations(digits string) []string {
	defer timeCost()()
	tempList, result, mp := make([][]byte, 0), make([]string, 0), map[rune][]byte{
		'2': {'a', 'b', 'c'},
		'3': {'d', 'e', 'f'},
		'4': {'g', 'h', 'i'},
		'5': {'j', 'k', 'l'},
		'6': {'m', 'n', 'o'},
		'7': {'p', 'q', 'r', 's'},
		'8': {'t', 'u', 'v'},
		'9': {'w', 'x', 'y', 'z'},
	}

	for _, s := range digits {
		v, _ := mp[s]
		tempList = append(tempList, v)
	}

	for _, list := range tempList {
		if len(result) == 0 {
			for _, item := range list {
				result = append(result, string(item))
			}
		} else {
			temp := make([]string, 0)
			for _, item := range list {
				for _, r := range result {
					temp = append(temp, r+string(item))
				}
			}
			result = temp
		}
	}
	return result
}
