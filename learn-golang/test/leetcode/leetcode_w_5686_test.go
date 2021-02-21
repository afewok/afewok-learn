package leetcode

import (
	"fmt"
	"testing"
)

/**
 * 5686. 移动所有球到每个盒子所需的最小操作数
 */

func Test_leetcode_5686(t *testing.T) {
	fmt.Println(minOperations("110"))
	fmt.Println(minOperations("001011"))
}

func minOperations(boxes string) []int {
	defer timeCost()()
	mp, count, result := make(map[int]bool), 0, make([]int, 0)
	for i, v := range boxes {
		if v == '1' {
			mp[i] = true
		}
	}
	for i := range boxes {
		count = 0
		for k := range mp {
			if i == k {
				continue
			}
			temp := k - i
			if temp < 0 {
				temp = -temp
			}
			count += temp
		}
		result = append(result, count)
	}
	return result
}
