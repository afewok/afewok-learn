package leetcode

import (
	"fmt"
	"testing"
)

/**
 * 188. 买卖股票的最佳时机 IV
 */

func Test_leetcode_188(t *testing.T) {
	fmt.Println(maxProfit(2, []int{2, 4, 1}))
	fmt.Println(maxProfit(2, []int{3, 2, 6, 5, 0, 3}))
}

func maxProfit(k int, prices []int) int {
	defer timeCost()()
	length := len(prices)
	if length < 1 {
		return 0
	}
	//当k非常大时转为无限次交易
	if k >= length/2 {
		dp0, dp1 := 0, -prices[0]
		for i := 1; i < length; i++ {
			dp0 = maxTwo(dp0, dp1+prices[i])
			dp1 = maxTwo(dp1, dp0-prices[i])
		}
		return maxTwo(dp0, dp1)
	}
	//定义二维数组，交易了多少次、当前的买卖状态
	dp := make([][]int, k+1)
	for i := 0; i < k+1; i++ {
		dp[i] = make([]int, 2)
	}
	for i := 0; i <= k; i++ {
		dp[i][0] = 0
		dp[i][1] = -prices[0]
	}
	for i := 1; i < length; i++ {
		for j := k; j > 0; j-- {
			//处理第k次买入
			dp[j-1][1] = maxTwo(dp[j-1][1], dp[j-1][0]-prices[i])
			//处理第k次卖出
			dp[j][0] = maxTwo(dp[j][0], dp[j-1][1]+prices[i])
		}
	}
	return dp[k][0]
}
