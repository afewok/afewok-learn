package leetcode

import (
	"fmt"
	"sort"
	"testing"
)

/**
 * 1202. 交换字符串中的元素
 */

func Test_leetcode_1202(t *testing.T) {
	fmt.Println(smallestStringWithSwaps("dcab", [][]int{{0, 3}, {1, 2}}))
	fmt.Println(smallestStringWithSwaps("dcab", [][]int{{0, 3}, {1, 2}, {0, 2}}))
	fmt.Println(smallestStringWithSwaps("cab", [][]int{{0, 1}, {1, 2}}))
}

func smallestStringWithSwaps(s string, pairs [][]int) string {
	defer timeCost()()
	if len(pairs) == 0 {
		return s
	}
	m := make(map[int][]int, 0)
	a := make([]int, len(s))
	for i := 0; i < len(s); i++ {
		a[i] = i
	}
	for _, pair := range pairs {
		merge1202(pair[0], pair[1], a)
	}
	for i := 0; i < len(s); i++ {
		t := getF1202(a, i)
		m[t] = append(m[t], i)
	}
	ans := make([]byte, len(s))
	for _, v := range m {
		tmp := make([]byte, len(v))
		for i := 0; i < len(v); i++ {
			tmp[i] = s[v[i]]
		}
		sort.Slice(tmp, func(i, j int) bool {
			return tmp[i] < tmp[j]
		})
		for i := 0; i < len(v); i++ {
			ans[v[i]] = tmp[i]
		}
	}
	return string(ans)
}

func merge1202(t1, t2 int, a []int) {
	t1 = getF1202(a, t1)
	t2 = getF1202(a, t2)
	if t1 != t2 {
		a[t2] = t1
	}
}

func getF1202(a []int, v int) int {
	if a[v] == v {
		return v
	} else {
		a[v] = getF1202(a, a[v])
		return a[v]
	}
}
