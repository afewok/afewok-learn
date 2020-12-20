package leetcode

import (
	"fmt"
	"testing"
)

/**
 * 746. 使用最小花费爬楼梯
 *
 * 思路：
 */
func Test_leetcode_746(t *testing.T) {
	fmt.Println(minCostClimbingStairs([]int{0, 2, 2, 1}))
	fmt.Println(minCostClimbingStairs([]int{10, 15, 20}))
	fmt.Println(minCostClimbingStairs([]int{1, 100, 1, 1, 1, 100, 1, 1, 100, 1}))
}

func minCostClimbingStairs(cost []int) int {
	defer timeCost()()
	expenditure, length := 0, len(cost)
	for i := 1; i < length; i++ {
		if cost[i-1] > cost[i] {
			expenditure += cost[i]
		}else 
	}
	return expenditure
}
