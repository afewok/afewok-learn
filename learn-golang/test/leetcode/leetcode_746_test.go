package leetcode

import (
	"fmt"
	"math"
	"testing"
)

/**
 * 746. 使用最小花费爬楼梯
 *
 * 思路：
 */
func Test_leetcode_746(t *testing.T) {
	fmt.Println(minCostClimbingStairs([]int{0, 1, 2, 2}))
	fmt.Println(minCostClimbingStairs([]int{1, 0, 0, 1}))
	fmt.Println(minCostClimbingStairs([]int{1, 2, 2, 0}))
	fmt.Println(minCostClimbingStairs([]int{0, 2, 2, 1}))
	fmt.Println(minCostClimbingStairs([]int{10, 15, 20}))
	fmt.Println(minCostClimbingStairs([]int{1, 100, 1, 1, 1, 100, 1, 1, 100, 1}))
}

func minCostClimbingStairs(cost []int) int {
	dp := make([]int, len(cost))
	dp[0] = cost[0]
	dp[1] = cost[1]
	for i := 2; i < len(cost); i++ {
		dp[i] = int(math.Min(float64(dp[i-1]), float64(dp[i-2]))) + cost[i]
	}
	return int(math.Min(float64(dp[len(cost)-1]), float64(dp[len(cost)-2])))
}

func minCostClimbingStairs1(cost []int) int {
	defer timeCost()()
	return getExpenditure(cost, -1, len(cost))
}

func getExpenditure(cost []int, sub, length int) int {
	if length == sub {
		return 0
	} else if length == sub+1 {
		return cost[sub]
	}
	temp := 0
	if sub != -1 {
		temp = cost[sub]
	}

	result1, result2 := temp+getExpenditure(cost, sub+1, length), temp+getExpenditure(cost, sub+2, length)
	if result1 < result2 {
		return result1
	}
	return result2
}
