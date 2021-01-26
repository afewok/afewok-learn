package leetcode

import (
	"fmt"
	"testing"
)

/**
 * 1128. 等价多米诺骨牌对的数量
 */

func Test_leetcode_1128(t *testing.T) {
	fmt.Println(numEquivDominoPairs([][]int{{1, 2}, {2, 1}, {3, 4}, {5, 6}}))
}

func numEquivDominoPairs(dominoes [][]int) int {
	defer timeCost()()
	count, mp := 0, make(map[int]map[int]int)
	for _, dominoe := range dominoes {
		min := getMin(dominoe[0], dominoe[1])
		max := getMax(dominoe[0], dominoe[1])
		if m, OK1 := mp[min]; OK1 {
			m[max]++
		} else {
			mp[min] = map[int]int{max: 1}
		}
	}

	for _, v1 := range mp {
		for _, v2 := range v1 {
			count += (v2 * (v2 - 1)) / 2
		}
	}
	return count
}

func getMax(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func getMin(x, y int) int {
	if x < y {
		return x
	}
	return y
}
