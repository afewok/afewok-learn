package leetcode

import (
	"fmt"
	"math"
	"testing"
)

/**
 * 50. Pow(x, n)
 *
 * 思路：
 */
func Test_leetcode_050(t *testing.T) {
	fmt.Println(myPow(2.00000, 10))
	fmt.Println(myPow(2.10000, 3))
	fmt.Println(myPow(2.00000, -2))
}

func myPow(x float64, n int) float64 {
	defer timeCost()()
	return math.Pow(x, float64(n))
}
