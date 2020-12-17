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
	fmt.Println(maxProfit(prices, 2))
}

type Solutions struct {
	Sum int
	Buy int
}

func maxProfit(prices []int, fee int) int {
	defer timeCost()()
	length, solutionsList, sum := len(prices), make([]Solutions), 0
	for i := 0; i < length; i++ {
		for _, solutions := range solutionsList {
		}
	}

	for _, solutions := range solutionsList {
		if sum < solutions.Sum {
			sum = solutions.Sum
		}
	}

	return sum
}

// func getProfit(prices []int, fee, length, sum, buy, sub int) int {
// 	if sub == length {
// 		return sum
// 	}

// 	price := prices[sub]
// 	sub++
// 	if buy == -1 {
// 		getProfit(prices, fee, length, sum, price, sub)
// 		getProfit(prices, fee, length, sum, 0, sub)
// 	} else {
// 		sell := price - (buy + fee)
// 		if sell > 0 {
// 			getProfit(prices, fee, length, sum+sell, 0, sub)
// 			getProfit(prices, fee, length, sum, buy, sub)
// 		} else {
// 			getProfit(prices, fee, length, sum, buy, sub)
// 		}
// 	}
// }
