package testing

import (
	"fmt"
	"testing"
	"time"
)

func Test_fill_cpu(t *testing.T) {
	var num uint32 = 0
	fmt.Print("请输入执行时间（单位：分钟）：")
	fmt.Scanln(&num)

	title := fmt.Sprintf("满CPU执行%d分钟,开始....", num)
	fmt.Println(title)

	quit := make(chan bool)
	for i := 0; i < 100; i++ {
		go func() {
			for {
				select {
				case <-quit:
					break
				default:
				}
			}
		}()
	}

	time.Sleep(time.Minute * time.Duration(num))
	for i := 0; i != 5; i++ {
		quit <- true
	}
}
