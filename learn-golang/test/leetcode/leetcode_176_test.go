package leetcode

import (
	"fmt"
	"testing"
)

/**
 * 176. 第二高的薪水
 */

func Test_leetcode_176(t *testing.T) {
	fmt.Println("排序去重分页", "SELECT (SELECT DISTINCT Salary FROM Employee ORDER BY Salary DESC LIMIT 1 OFFSET 1) AS SecondHighestSalary")
	fmt.Println("排序去重分页+IFNULL", "SELECT IFNULL((SELECT DISTINCT Salary FROM Employee ORDER BY Salary DESC  LIMIT 1 OFFSET 1), NULL) AS SecondHighestSalary")
	fmt.Println("通过子查询去掉第一大的数,即可得到第二大的数", "select max(Salary) SecondHighestSalary  from Employee where Salary != (select max(Salary) from Employee)")
	fmt.Println("通过子查询去掉第一大的数,即可得到第二大的数", "select max(Salary) SecondHighestSalary from employee where salary<(select max(salary) from employee)")

	fmt.Println("", "select(select e1.Salary from Employee e1,Employee e2 where e1.Salary<=e2.Salary group by e1.Salary having count(distinct e2.Salary)=2) as SecondHighestSalary")
	fmt.Println("", "select(select distinct e1.Salary from Employee e1,Employee e2 group by e1.Id,e1.Salary having sum(e1.Salary<e2.Salary)=1 ) as SecondHighestSalary")
	fmt.Println("窗口函数", "SELECT IFNULL((SELECT DISTINCT Salary FROM (SELECT Salary, DENSE_RANK() OVER (ORDER BY Salary DESC) AS Num FROM Employee ) tb WHERE Num = 2), NULL) AS SecondHighestSalary")
}
