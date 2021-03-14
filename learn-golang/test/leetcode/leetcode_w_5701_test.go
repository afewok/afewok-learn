package leetcode

import (
	"fmt"
	"testing"
)

/**
 * 5701. 仅执行一次字符串交换能否使两个字符串相等
 */

func Test_leetcode_5701(t *testing.T) {
	// fmt.Println(areAlmostEqual("bank", "kanb"))
	// fmt.Println(areAlmostEqual("attack", "defend"))
	// fmt.Println(areAlmostEqual("kelb", "kelb"))
	// fmt.Println(areAlmostEqual("abcd", "dcba"))
	fmt.Println(areAlmostEqual("gasitigykqhrtnypjbeqwwsngclzr", "gasitigykqhrtnypjbeqwwsngczlr"))
}

func areAlmostEqual(s1 string, s2 string) bool {
	defer timeCost()()
	lenS1, lenS2 := len(s1), len(s2)
	if lenS1 != lenS2 {
		return false
	}

	arr1, arr2 := make([][]byte, 0), make([][]byte, 0)
	for i := 0; i < lenS1; i++ {
		if s1[i] != s2[i] {
			if i == 0 || s1[i-1] == s2[i-1] {
				arr1 = append(arr1, []byte{s1[i]})
				arr2 = append(arr2, []byte{s2[i]})
			} else {
				arr1[len(arr1)-1] = append(arr1[len(arr1)-1], s1[i])
				arr2[len(arr2)-1] = append(arr2[len(arr2)-1], s2[i])
			}
			if len(arr1) > 2 {
				return false
			}
		}
	}
	if len(arr1) == 0 {
		return true
	} else if len(arr1) == 1 && len(arr1[0]) == 2 && len(arr1[0]) == len(arr2[0]) {
		if arr1[0][0] == arr2[0][1] && arr1[0][1] == arr2[0][0] {
			return true
		}
		return false
	} else if len(arr1) != 2 || len(arr1[0]) != len(arr2[1]) || len(arr1[1]) != len(arr2[0]) {
		return false
	}
	for i, b := range arr1[0] {
		if b != arr2[1][i] {
			return false
		}
	}

	for i, b := range arr1[1] {
		if b != arr2[0][i] {
			return false
		}
	}
	return true
}
