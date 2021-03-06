package leetcode

import (
	"fmt"
	"testing"
)

/**
 * 577. 员工奖金
 */

func Test_leetcode_577(t *testing.T) {
	fmt.Println("left join", "select * from (select e.name as name , b.bonus as bonus from Employee e  left join Bonus b on e.empId =b.empId) as t where  t.bonus  <1000 or t.bonus is null")
}
