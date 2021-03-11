package leetcode

import (
	"fmt"
	"testing"
)

/**
 * 193. 有效电话号码
 */

func Test_leetcode_193(t *testing.T) {
	fmt.Println("grep", "grep -P '^([0-9]{3}-|\\([0-9]{3}\\) )[0-9]{3}-[0-9]{4}$' file.txt")
	fmt.Println("grep", "grep -P '^(\\d{3}-|\\(\\d{3}\\) )\\d{3}-\\d{4}$' file.txt")
	fmt.Println("awk", "awk '/^([0-9]{3}-|\\([0-9]{3}\\) )[0-9]{3}-[0-9]{4}$/' file.txt")
	fmt.Println("gawk", "'/^([0-9]{3}-|\\([0-9]{3}\\) )[0-9]{3}-[0-9]{4}$/' file.txt")
	fmt.Println("sed", "sed -n '/^[0-9]\\{3\\}-[0-9]\\{3\\}-[0-9]\\{4\\}$\\|^([0-9]\\{3\\}) [0-9]\\{3\\}-[0-9]\\{4\\}$/p' file.txt")
	fmt.Println("cat", "cat file.txt | grep -P '^(\\([0-9]{3}\\)\\s|[0-9]{3}-)[0-9]{3}-[0-9]{4}$'")
}
