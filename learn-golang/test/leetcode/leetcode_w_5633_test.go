package leetcode

import (
	"fmt"
	"testing"
)

/**
 * 5633. 计算力扣银行的钱
 */

func Test_leetcode_5633(t *testing.T) {
	fmt.Println(totalMoney(4))
	fmt.Println(totalMoney(10))
	fmt.Println(totalMoney(20))
}

func totalMoney(n int) int {
	defer timeCost()()
	cycle, surplus, result := n/7, n%7, 0

	for surplus > 0 {
		result = result + surplus + cycle
		surplus--
	}
	cycle--
	for cycle >= 0 {
		result = result + 28 + cycle*7
		cycle--
	}
	return result
}
