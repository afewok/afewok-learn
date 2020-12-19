package leetcode

import (
	"fmt"
	"testing"
)

/**
 * 121. 买卖股票的最佳时机
 */

func Test_leetcode_121(t *testing.T) {
	fmt.Println(maxProfit121([]int{7, 1, 5, 3, 6, 4}))
}

func maxProfit121(prices []int) int {
	defer timeCost()()
	length := len(prices)
	if length < 2 {
		return 0
	}
	buy, sale, temp, sum := prices[0], prices[0], 0, 0
	for i := 1; i < length; i++ {
		if buy > prices[i] {
			temp = sale - buy
			if sum < temp {
				sum = temp
			}
			buy = prices[i]
			sale = prices[i]
		} else if sale < prices[i] {
			sale = prices[i]
		}
	}
	temp = sale - buy
	if sum < temp {
		sum = temp
	}
	return sum
}
