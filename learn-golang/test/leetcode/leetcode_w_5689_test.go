package leetcode

import (
	"fmt"
	"testing"
)

/**
 * 5689. 统计匹配检索规则的物品数量
 */

func Test_leetcode_5689(t *testing.T) {
	fmt.Println(countMatches([][]string{{"phone", "blue", "pixel"}, {"computer", "silver", "lenovo"}, {"phone", "gold", "iphone"}}, "color", "silver"))
	fmt.Println(countMatches([][]string{{"phone", "blue", "pixel"}, {"computer", "silver", "phone"}, {"phone", "gold", "iphone"}}, "type", "phone"))
}

func countMatches(items [][]string, ruleKey string, ruleValue string) int {
	defer timeCost()()
	mp, result := map[string]int{"type": 0, "color": 1, "name": 2}, 0
	sub, OK := mp[ruleKey]
	if !OK {
		return result
	}
	for _, item := range items {
		if item[sub] == ruleValue {
			result++
		}
	}
	return result
}
