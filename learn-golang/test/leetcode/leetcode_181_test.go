package leetcode

import (
	"fmt"
	"testing"
)

/**
 * 181. 超过经理收入的员工
 */

func Test_leetcode_181(t *testing.T) {
	fmt.Println("left join...on", "select e1.Name as Employee from Employee e1 left join Employee e2 on e1.ManagerId = e2.Id where e1.Salary >e2.Salary")
	fmt.Println("join (可省略)", "select e1.Name as Employee from Employee e1, Employee e2 where e1.ManagerId = e2.Id and e1.Salary >e2.Salary")
	fmt.Println("", "select e.Name as Employee from employee e where salary > (select salary from employee where Id = e.ManagerId)")
}
