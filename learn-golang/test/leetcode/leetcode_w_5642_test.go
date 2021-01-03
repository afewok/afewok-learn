package leetcode

import (
	"fmt"
	"testing"
)

/**
 * 5642. 大餐计数
 */

func Test_leetcode_5642(t *testing.T) {
	fmt.Println(countPairs([]int{1, 3, 5, 7, 9}))
	fmt.Println(countPairs([]int{1, 1, 1, 3, 3, 3, 7}))
	fmt.Println(countPairs([]int{149, 107, 1, 63, 0, 1, 6867, 1325, 5611, 2581, 39, 89, 46, 18, 12, 20, 22, 234}))
}

func countPairs(deliciousness []int) int {
	defer timeCost()()
	MOD, res, cnt, k, m := 1_000_000_000+7, 0, make([]int, (1<<20)+1), 1<<21, 1<<20

	for _, d := range deliciousness {
		for x := 1; x <= k; x *= 2 {
			if 0 <= x-d && x-d <= m {
				res = (res + cnt[x-d]) % MOD
			}
		}
		cnt[d]++
	}
	return int(res)
}
