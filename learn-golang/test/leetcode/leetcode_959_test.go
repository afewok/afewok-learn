package leetcode

import (
	"fmt"
	"testing"
)

/**
 * 959. 由斜杠划分区域
 */

func Test_leetcode_959(t *testing.T) {
	fmt.Println(regionsBySlashes([]string{" /", "/ "}))
}

func regionsBySlashes(grid []string) int {
	defer timeCost()()
	rmap := make(map[int]int, 0)
	ll := len(grid)
	s := NewSet(4 * ll * ll)

	for i := 0; i < ll; i++ {

		for j := 0; j < ll; j++ {
			zero := i*4*ll + 4*j
			switch grid[i][j] {
			case '\\':
				s.Union(zero+1, zero+2)
				s.Union(zero+0, zero+3)
			case '/':
				s.Union(zero+0, zero+1)
				s.Union(zero+2, zero+3)
			case ' ':
				s.Union(zero+0, zero+1)
				s.Union(zero+0, zero+2)
				s.Union(zero+0, zero+3)
			}

			if j+1 != ll {
				s.Union(zero+2, zero+4)
			}
			if i+1 != ll {
				s.Union(zero+3, zero+4*ll+1)
			}
		}
	}

	for i := 0; i < 4*ll*ll; i++ {
		rmap[s.Find(i)] = 0
	}

	return len(rmap)
}

type Set struct {
	s []int
}

func NewSet(n int) *Set {
	s := make([]int, n)
	for i := 0; i < len(s); i++ {
		s[i] = i
	}
	return &Set{
		s,
	}
}

func (s *Set) Union(a, b int) {
	if s.Find(a) == s.Find(b) {
		return
	} else {
		s.s[s.Find(b)] = s.s[s.Find(a)]
	}
}

func (s *Set) Find(x int) int {
	if s.s[x] == x {
		return x
	} else {
		return s.Find(s.s[x])
	}
}
