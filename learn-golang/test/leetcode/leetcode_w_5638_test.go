package leetcode

import (
	"fmt"
	"testing"
)

/**
 * 5638. 吃苹果的最大数目
 */

func Test_leetcode_5638(t *testing.T) {
	fmt.Println(eatenApples([]int{1, 2, 3, 5, 2}, []int{3, 2, 1, 4, 2}))
	fmt.Println(eatenApples([]int{3, 0, 0, 0, 0, 2}, []int{3, 0, 0, 0, 0, 2}))
}

func eatenApples(apples []int, days []int) int {
	defer timeCost()()
	length, max, eat := len(apples), 0, 0
	list := make([]int, 20000+length)
	for i, item := range list {
		if i < length {
			temp := days[i] + i
			list[temp] = apples[i] + item
			if max < temp {
				max = temp
			}
		} else if i >= max {
			break
		}
		for j := i + 1; j <= max; j++ {
			if list[j] > 0 {
				list[j]--
				eat++
				break
			}
		}
	}
	return eat
}
