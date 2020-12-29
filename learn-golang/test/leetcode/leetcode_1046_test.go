package leetcode

import (
	"fmt"
	"testing"
)

/**
 * 1046. 最后一块石头的重量
 */

func Test_leetcode_1046(t *testing.T) {
	fmt.Println(lastStoneWeight([]int{2, 7, 4, 1, 8, 1}))
	fmt.Println(lastStoneWeight([]int{2}))
	fmt.Println(lastStoneWeight([]int{2, 2}))
}

func lastStoneWeight(stones []int) int {
	defer timeCost()()
	sub, length := 0, len(stones)
	for sub+1 < length {
		for i := sub; i < sub+2; i++ {
			for j := i + 1; j < length; j++ {
				if stones[i] < stones[j] {
					stones[i], stones[j] = stones[j], stones[i]
				}
			}
		}

		if stones[sub] == 0 {
			return stones[sub]
		} else if stones[sub] == stones[sub+1] {
			sub = sub + 2
		} else {
			stones[sub+1] = stones[sub] - stones[sub+1]
			sub = sub + 1
		}
	}
	if sub >= length {
		return 0
	}
	return stones[sub]
}
