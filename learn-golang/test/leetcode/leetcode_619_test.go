package leetcode

import (
	"fmt"
	"testing"
)

/**
 * 619. 只出现一次的最大数字
 */

func Test_leetcode_619(t *testing.T) {
	fmt.Println("IFNULL  group by ... having   order by ... DESC LIMIT 1 OFFSET 0", "SELECT IFNULL((select num from my_numbers group by num having count(1)=1 order by num DESC  LIMIT 1 OFFSET 0),NULL) as num")
}
