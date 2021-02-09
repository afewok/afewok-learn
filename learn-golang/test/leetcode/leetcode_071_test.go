package leetcode

import (
	"fmt"
	"strings"
	"testing"
)

/**
 * 71. 简化路径
 */

func Test_leetcode_071(t *testing.T) {
	fmt.Println(simplifyPath("home/"))
	fmt.Println(simplifyPath("/home/"))
	fmt.Println(simplifyPath("/../"))
	fmt.Println(simplifyPath("/home//foo/"))
	fmt.Println(simplifyPath("/a/./b/../../c/"))
}

func simplifyPath(path string) string {
	defer timeCost()()
	result, strs := make([]string, 0), strings.Split(path, "/")
	if strs[0] != "" {
		return "/"
	}
	for _, str := range strs {
		if str == "" || str == "." {
			continue
		} else if str == ".." {
			length := len(result)
			if length > 0 {
				result = result[0 : length-1]
			}
		} else {
			result = append(result, str)
		}
	}
	return "/" + strings.Join(result, "/")
}
