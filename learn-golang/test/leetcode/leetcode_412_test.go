package leetcode

import (
	"fmt"
	"strconv"
	"testing"
)

/**
 * 412. Fizz Buzz
 */
func Test_leetcode_412(t *testing.T) {
	fmt.Println(fizzBuzz(15))
}

func fizzBuzz(n int) []string {
	defer timeCost()()
	result := make([]string, n)
	for i := 1; i <= n; i++ {
		temp1, temp2 := i%3 == 0, i%5 == 0
		if temp1 && temp2 {
			result[i-1] = "FizzBuzz"
		} else if temp1 {
			result[i-1] = "Fizz"
		} else if temp2 {
			result[i-1] = "Buzz"
		} else {
			result[i-1] = strconv.Itoa(i)
		}
	}
	return result
}
