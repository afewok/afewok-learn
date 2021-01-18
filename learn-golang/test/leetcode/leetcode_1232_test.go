package leetcode

import (
	"fmt"
	"testing"
)

/**
 * 1232. 缀点成线
 */

func Test_leetcode_1232(t *testing.T) {
	fmt.Println(checkStraightLine([][]int{{1, 2}, {2, 3}, {3, 4}, {4, 5}, {5, 6}, {6, 7}}))
	fmt.Println(checkStraightLine([][]int{{1, 1}, {2, 2}, {3, 4}, {4, 5}, {5, 6}, {7, 7}}))
	fmt.Println(checkStraightLine([][]int{{0, 0}, {0, 1}, {0, -1}}))
}

func checkStraightLine(coordinates [][]int) bool {
	defer timeCost()()
	deltaX, deltaY := coordinates[0][0], coordinates[0][1]
	for _, p := range coordinates {
		p[0] -= deltaX
		p[1] -= deltaY
	}
	A, B := coordinates[1][1], -coordinates[1][0]
	for _, p := range coordinates[2:] {
		x, y := p[0], p[1]
		if A*x+B*y != 0 {
			return false
		}
	}
	return true
}
