package leetcode

import (
	"fmt"
	"testing"
)

/**
 * 182. 查找重复的电子邮箱
 */

func Test_leetcode_182(t *testing.T) {
	fmt.Println("GROUP BY column_name HAVING aggregate_function(column_name) operator value", "select Email from Person group by Email having count(Email)>1")
	fmt.Println("", "select email from (select count(1) as t,email from person group by email)r  where r.t>1")
	fmt.Println("", "select distinct(p1.Email) from Person p1 join Person  p2 on p1.Email = p2.Email AND p1.Id!=p2.Id")
}
