package leetcode

import (
	"fmt"
	"testing"
)

/**
 * 135. 分发糖果
 */

func Test_leetcode_135(t *testing.T) {
	// fmt.Println(candy([]int{1, 0, 2}))
	// fmt.Println(candy([]int{1, 2, 2}))
	// fmt.Println(candy([]int{1, 3, 2, 2, 1}))
	// fmt.Println(candy([]int{1, 2, 87, 87, 87, 2, 1}))
	// fmt.Println(candy([]int{29, 51, 87, 87, 72, 12}))
	// fmt.Println(candy([]int{1, 6, 10, 8, 7, 3, 2}))
	fmt.Println(candy([]int{1, 3, 4, 5, 2}))
}

// 5+1+2+3
func candy(ratings []int) int {
	defer timeCost()()
	length, award, i := len(ratings), 0, 1
	arr := make([]int, length)

	for i < length {
		if ratings[i-1] > ratings[i] {
			j := i + 1
			for j < length {
				if ratings[j-1] <= ratings[j] {
					break
				}
				j++
			}
			count := 1
			for k := j - 2; k >= i; k-- {
				arr[k] = count
				award += arr[k]
				count++
			}
			if i-1 == 0 || ratings[i-1] == ratings[i-2] {
				arr[i-1] = count
				award += arr[i-1]
			} else if ratings[i-1] > ratings[i-2] {
				if count > arr[i-1] {
					award += count - arr[i-1]
				}
				arr[i-1] = count
			}
			i = j - 1
		} else if ratings[i-1] < ratings[i] {
			arr[i] = arr[i-1] + 1
			award += arr[i]
		}
		i++
	}
	return length + award
}
