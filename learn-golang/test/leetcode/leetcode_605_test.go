package leetcode

import (
	"fmt"
	"testing"
)

/**
 * 605. 种花问题
 */

func Test_leetcode_605(t *testing.T) {
	// fmt.Println(canPlaceFlowers([]int{1, 0, 0, 0, 1}, 1))
	// fmt.Println(canPlaceFlowers([]int{1, 0, 0, 0, 1}, 2))
	fmt.Println(canPlaceFlowers([]int{1, 0, 0, 0, 0, 1}, 2))
}

func canPlaceFlowers(flowerbed []int, n int) bool {
	defer timeCost()()
	if n == 0 {
		return true
	}
	length := len(flowerbed)
	max := length - 1
	for i, v := range flowerbed {
		if v == 1 {
			continue
		}

		if i == 0 {
			if length == 1 || flowerbed[i+1] == 0 {
				flowerbed[i] = 1
				n--
			}
		} else if i == max {
			if flowerbed[i-1] == 0 {
				flowerbed[i] = 1
				n--
			}
		} else if flowerbed[i-1] == 0 && flowerbed[i+1] == 0 {
			flowerbed[i] = 1
			n--
		}

		if n == 0 {
			return true
		}
	}
	return false
}
