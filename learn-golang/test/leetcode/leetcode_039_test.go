package leetcode

import (
	"fmt"
	"testing"
)

/**
 * 39. 组合总和
 */

func Test_leetcode_039(t *testing.T) {
	fmt.Println(combinationSum([]int{2, 3, 6, 7}, 7))
	fmt.Println(combinationSum([]int{2, 3, 5}, 8))
}

func combinationSum(candidates []int, target int) [][]int {
	defer timeCost()()
	mp, length, result := make(map[int]bool), len(candidates), make([][]int, 0)
	for i := 0; i < length; i++ {
		mp[candidates[i]] = true
	}
	for i := 1; i <= target && i <= length; i++ {
		temp := getCombinationSum(mp, i, target)
		if temp != nil {
			result = append(result, temp...)
		}
	}
	return result
}

func getCombinationSum(mp map[int]bool, count, target int) [][]int {

}
