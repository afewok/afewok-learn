package leetcode

import (
	"fmt"
	"testing"
)

/**
 * 46. 全排列
 */

func Test_leetcode_046(t *testing.T) {
	fmt.Println(permute([]int{5, 4, 6, 2}))
	fmt.Println(permute([]int{1, 2, 3}))
}

func permute(nums []int) [][]int {
	defer timeCost()()
	length := len(nums)
	result := make([][]int, 0)
	for i := 0; i < length; i++ {
		if i == 0 {
			result = append(result, []int{nums[0]})
		} else {
			lenR := len(result)
			for m := 0; m < lenR; m++ {
				for n := 0; n < len(result[m]); n++ {
					temp := make([]int, len(result[m])+1)
					copy(temp[:n], result[m][:n])
					copy(temp[n+1:], result[m][n:])
					temp[n] = nums[i]
					result = append(result, temp)
				}
			}
			for j := 0; j < lenR; j++ {
				result[j] = append(result[j], nums[i])
			}
		}
	}
	return result
}
