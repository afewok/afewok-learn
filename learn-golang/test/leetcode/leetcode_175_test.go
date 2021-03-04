package leetcode

import (
	"fmt"
	"testing"
)

/**
 * 175. 组合两个表
 */

func Test_leetcode_175(t *testing.T) {
	fmt.Println("多次查询", "select FirstName, LastName, (select City from Address a where a.PersonId=p.PersonId) as City, (select State from Address a where a.PersonId=p.PersonId) as State  from Person p")
	fmt.Println("left join ... on ", "select FirstName, LastName, City, State from Person left join Address on Person.PersonId = Address.PersonId")
	fmt.Println("left join ... on ", "select FirstName, LastName, City, State from Person left join Address USING(PersonId)")
	fmt.Println("left join ... on ", "select p.FirstName, p.LastName, a.City, a.State from Person p left join Address a on p.PersonId = a.PersonId")
}
