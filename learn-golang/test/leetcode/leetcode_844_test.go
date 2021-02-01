package leetcode

import (
	"fmt"
	"testing"
)

/**
 * 844. 比较含退格的字符串
 *
 */
func Test_leetcode_844(t *testing.T) {
	fmt.Println(backspaceCompare("ab#c", "ad#c"))
	fmt.Println(backspaceCompare("ab##", "c#d#"))
	fmt.Println(backspaceCompare("a##c", "#a#c"))
	fmt.Println(backspaceCompare("a#c", "b"))
}

func backspaceCompare(S string, T string) bool {
	defer timeCost()()
	tempSList, tempTList, tempSLen, tempTLen := make([]byte, len(S)), make([]byte, len(T)), 0, 0
	for _, v := range S {
		if v == '#' {
			if tempSLen > 0 {
				tempSLen--
			}
		} else {
			tempSList[tempSLen] = byte(v)
			tempSLen++
		}
	}

	for _, v := range T {
		if v == '#' {
			if tempTLen > 0 {
				tempTLen--
			}
		} else {
			tempTList[tempTLen] = byte(v)
			tempTLen++
		}
	}
	if tempSLen != tempTLen {
		return false
	}
	for i := 0; i < tempSLen; i++ {
		if tempSList[i] != tempTList[i] {
			return false
		}
	}

	return true
}
