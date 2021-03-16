package leetcode

import (
	"fmt"
	"testing"
)

/**
 * 47. 全排列 II
 */

func Test_leetcode_047(t *testing.T) {
	fmt.Println(permuteUnique([]int{1, 1, 2}))
	fmt.Println(permuteUnique([]int{1, 2, 3}))
}

func permuteUnique(nums []int) [][]int {
	defer timeCost()()
	length := len(nums)
	result := make([][]int, 0)
	for i := 0; i < length; i++ {
		if i == 0 {
			result = append(result, []int{nums[0]})
		} else {
			lenR := len(result)
			for j := 0; j < lenR; j++ {
				result[j] = append(result[j], nums[i])
			}
			for m := 0; m < lenR; m++ {
				for n := 0; n < len(result[m])-1; n++ {
					temp := make([]int, len(result[m]))
					copy(temp[:n], result[m][:n])
					copy(temp[n+1:], result[m][n:len(result[m])-1])
					temp[n] = nums[i]
					OK1 := true
					for q := 0; q < len(result); q++ {
						OK2 := true
						for p := 0; p < len(result[m]); p++ {
							if result[q][p] != temp[p] {
								OK2 = false
							}
						}
						if OK2 {
							OK1 = false
							break
						}
					}
					if OK1 {
						result = append(result, temp)
					}
				}
			}

		}
	}
	return result
}
