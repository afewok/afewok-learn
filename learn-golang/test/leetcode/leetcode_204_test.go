package leetcode

import (
	"fmt"
	"math"
	"testing"
)

/**
 * 204. 计数质数
 */

func Test_leetcode_204(t *testing.T) {
	fmt.Println(countPrimes(100))
	fmt.Println(countPrimes(5 * 1000000))
	fmt.Println(countPrimes(1500000))
}

func countPrimes1(n int) int {
	defer timeCost()()
	var chList []chan int
	sum, i := 0, 2
	if n > 10002 {
		for ; i < n; i += 10000 {
			tempCh := make(chan int)
			chList = append(chList, tempCh)
			go getPrimes(i, i+10000, tempCh)
		}
		tempCh := make(chan int)
		chList = append(chList, tempCh)
		go getPrimes(i-10000, n, tempCh)
	} else {
		tempCh := make(chan int)
		chList = append(chList, tempCh)
		go getPrimes(i, n, tempCh)
	}
	for _, ch := range chList {
		for {
			if -1 == <-ch {
				close(ch)
				break
			}
			sum++
		}
	}
	return sum
}

func getPrimes(start, end int, ch chan int) {
	temp, sub := 0, 0
	for i := start; i < end; i++ {
		temp, sub = sqrt(i), 2
		for ; sub <= temp; sub++ {
			if i%sub == 0 {
				break
			}
		}
		if sub > temp {
			ch <- 1
		}
	}
	ch <- -1
}

//牛顿法求平方根
func sqrt(x int) int {
	res := x
	for res*res > x {
		res = (res + x/res) / 2
	}
	return res
}

func countPrimes(n int) int {
	defer timeCost()()
	num := 0
	for i := 2; i < n; i++ {
		if isPrime(i) {
			num++
		}
	}
	return num
}

func isPrime(num int) bool {
	if num == 2 || num == 3 {
		return true
	}
	if num%6 != 1 && num%6 != 5 {
		return false
	}
	length := int(math.Sqrt(float64(num)))
	for i := 5; i <= length; i += 6 {
		if num%i == 0 || num%(i+2) == 0 {
			return false
		}
	}
	return true
}
