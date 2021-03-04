package leetcode

import (
	"fmt"
	"testing"
)

/**
 * 596. 超过5名学生的课
 */

func Test_leetcode_596(t *testing.T) {
	fmt.Println("三层查询", "select c.class from (select b.class,count(b.class) as nums from (select DISTINCT student, class from courses) as b group by b.class) as c where c.nums>=5")
	fmt.Println("二层查询", "select b.class from (select DISTINCT student, class from courses) as b group by b.class having count(b.class)>=5")
	fmt.Println("一层查询", "select class from courses group by class having count(DISTINCT STUDENT)>=5")
	fmt.Println("一层查询", "SELECT  class FROM (SELECT class, COUNT(DISTINCT student) AS num FROM courses GROUP BY class) AS temp_table WHERE num >= 5")
}
