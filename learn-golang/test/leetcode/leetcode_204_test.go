package leetcode

import (
	"fmt"
	"strconv"
	"strings"
	"testing"
)

/**
 * 204. 计数质数
 */

func Test_leetcode_204(t *testing.T) {
	// fmt.Println(countPrimes(100))
	fmt.Println(countPrimes(5 * 1000000))
	// fmt.Println(countPrimes(5 * 1500000))
}

func countPrimes(n int) int {
	defer timeCost()()
	for i, item := range arr {
		if n < item {
			return i
		}
	}
	return len(arr)
}

func countPrimes1(n int) int {
	defer timeCost()()
	var builder strings.Builder
	num, temp, sub := 0, 0, 0
	for i := 2; i < n; i++ {
		temp, sub = sqrt(i), 2
		for ; sub <= temp; sub++ {
			if i%sub == 0 {
				break
			}
		}
		if sub > temp {
			builder.WriteString(strconv.Itoa(i))
			builder.WriteString(",")
			num++
		}
	}
	fmt.Println(builder.String())
	return num
}

//牛顿法求平方根
func sqrt(x int) int {
	res := x
	for res*res > x {
		res = (res + x/res) / 2
	}
	return res
}