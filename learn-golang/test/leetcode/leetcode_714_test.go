package leetcode

import (
	"fmt"
	"testing"
)

/**
 * 714. 买卖股票的最佳时机含手续费
 *
 * 思路：
 */
func Test_leetcode_714(t *testing.T) {
	prices := []int{1, 3, 2, 8, 4, 9}
	fmt.Println(maxProfit714(prices, 2))
}

func maxProfit714(prices []int, fee int) int {
	defer timeCost()()
	n := len(prices)
	buy := prices[0] + fee
	profit := 0
	for i := 1; i < n; i++ {
		if prices[i]+fee < buy {
			buy = prices[i] + fee
		} else if prices[i] > buy {
			profit += prices[i] - buy
			buy = prices[i]
		}
	}
	return profit
}
