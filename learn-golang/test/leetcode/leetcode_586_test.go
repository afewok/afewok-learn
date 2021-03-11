package leetcode

import (
	"fmt"
	"testing"
)

/**
 * 586. 订单最多的客户
 */

func Test_leetcode_586(t *testing.T) {
	fmt.Println("distinct  group by   order by  desc ", "select o.customer_number as customer_number from (select distinct customer_number,count(customer_number) as num from orders group by customer_number) o order by o.num desc limit 1")
	fmt.Println("", "select customer_number from orders group by customer_number order by count(customer_number) desc limit 1")
	fmt.Println("", "select customer_number from orders group by customer_number  having count(*)=(select count(*) as c from orders group by customer_number order by c desc limit 1)")
}
