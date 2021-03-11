package leetcode

import (
	"fmt"
	"testing"
)

/**
 * 183. 从不订购的客户
 */

func Test_leetcode_183(t *testing.T) {
	fmt.Println("not in", "select c.Name  as Customers from Customers c where c.Id not in(select CustomerId  from Orders)")
	fmt.Println("NOT EXISTS", "select c.Name  as Customers from Customers c where NOT EXISTS(select 1  from Orders o where c.Id =o.CustomerId)")
	fmt.Println("", "SELECT t1.Name Customers FROM Customers t1 LEFT JOIN Orders t2 ON t1.Id = t2.CustomerId WHERE t2.CustomerId is null")
}
