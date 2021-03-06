package leetcode

import (
	"fmt"
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
	times := n
	//如果说n小于零，那么就求x的倒数的正n次方
	if times >= 0 {
		return qucikMulit(x, times)
	}
	return 1.0 / qucikMulit(x, -times)
}
func qucikMulit(x float64, times int) float64 {
	res := 1.0
	for times > 0 {
		if times%2 == 1 {
			res *= x
		}
		x = x * x
		times /= 2
	}
	return res
}
