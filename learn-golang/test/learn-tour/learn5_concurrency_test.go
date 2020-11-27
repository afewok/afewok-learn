package learntour

import (
	"fmt"
	"testing"
	"time"
)

//Go 程
func Test_goroutines(t *testing.T) {
	println("Go 程（goroutine）是由 Go 运行时管理的轻量级线程。")
	println("Go 程在相同的地址空间中运行，因此在访问共享的内存时必须进行同步。sync 包提供了这种能力，不过在 Go 中并不经常用到，因为还有其它的办法")
	go say("world")
	say("hello")
}

func say(s string) {
	for i := 0; i < 5; i++ {
		time.Sleep(1000 * time.Millisecond)
		fmt.Println(s)
	}
}

//信道
func Test_channels(t *testing.T) {
	println("信道是带有类型的管道，你可以通过它用信道操作符 <- 来发送或者接收值。")
	println("默认情况下，发送和接收操作在另一端准备好之前都会阻塞。这使得 Go 程可以在没有显式的锁或竞态变量的情况下进行同步。")
	s := []int{7, 2, 8, -9, 4, 0}
	c := make(chan int)
	go sum(s[:len(s)/2], c)
	go sum(s[len(s)/2:], c)
	x, y := <-c, <-c //从c中接受

	fmt.Println(x, y, x+y)

}

func sum(s []int, c chan int) {
	sum := 0
	for _, v := range s {
		sum += v
	}
	c <- sum //将和送入c
}

//带缓冲的信道
func Test_buffered_channels(t *testing.T) {
	println("信道可以是 带缓冲的。将缓冲长度作为第二个参数提供给 make 来初始化一个带缓冲的信道：")
	println("仅当信道的缓冲区填满后，向其发送数据时才会阻塞。当缓冲区为空时，接受方会阻塞。")
	ch := make(chan int, 2)
	ch <- 1
	ch <- 2
	fmt.Println(<-ch)
	fmt.Println(<-ch)
}

//range 和 close
func Test_range_and_close(t *testing.T) {
	println("发送者可通过 close 关闭一个信道来表示没有需要发送的值了")
	println("接收者可以通过为接收表达式分配第二个参数来测试信道是否被关闭 v, ok := <-ch")
	println("循环 for i := range c 会不断从信道接收值，直到它被关闭")
	println("*注意：* 只有发送者才能关闭信道，而接收者不能。向一个已经关闭的信道发送数据会引发程序恐慌（panic）")
	println("*还要注意：* 信道与文件不同，通常情况下无需关闭它们。只有在必须告诉接收者不再有需要发送的值时才有必要关闭，例如终止一个 range 循环。")
	c := make(chan int, 10)
	go fibonacci2(cap(c), c)
	for i := range c {
		fmt.Println(i)
	}
}

func fibonacci2(n int, c chan int) {
	x, y := 0, 1
	for i := 0; i < n; i++ {
		c <- x
		x, y = y, x+y
	}
	close(c)
}

//select 语句
func Test_select(t *testing.T) {

}

//默认选择
func Test_default_selection(t *testing.T) {

}

//练习：等价二叉查找树
func Test_exercise_equivalent_binary_trees(t *testing.T) {
}

//sync.Mutex
func Test_mutex_counter(t *testing.T) {
}

//练习：Web 爬虫
func Test_mutex_counter(t *testing.T) {
}
