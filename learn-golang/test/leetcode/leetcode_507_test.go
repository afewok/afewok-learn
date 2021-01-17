package leetcode

import (
	"fmt"
	"testing"
)

/**
 * 507. 完美数
 */

func Test_leetcode_507(t *testing.T) {
	fmt.Println(checkPerfectNumber(28))
	fmt.Println(checkPerfectNumber(6))
	fmt.Println(checkPerfectNumber(496))
	fmt.Println(checkPerfectNumber(8128))
	fmt.Println(checkPerfectNumber(2))
}

func checkPerfectNumber(num int) bool {
	defer timeCost()()
	primes := []int{2, 3, 5, 7, 13, 17, 19, 31}
	for _, p := range primes {
		if (1<<(p-1))*((1<<p)-1) == num {
			return true
		}
	}
	return false
}
