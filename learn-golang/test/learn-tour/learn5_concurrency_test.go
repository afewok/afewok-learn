package learntour

import (
	"fmt"
	"sync"
	"testing"
	"time"
	// "golang.org/x/tour/tree"
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
	println("select 语句使一个 Go 程可以等待多个通信操作。")
	println("select 会阻塞到某个分支可以继续执行为止，这时就会执行该分支。当多个分支都准备好时会随机选择一个执行。")
	c := make(chan int)
	quit := make(chan int)
	go func() {
		for i := 0; i < 10; i++ {
			fmt.Println(<-c)
		}
		quit <- 0
	}()
	fibonacci3(c, quit)
}

func fibonacci3(c, quit chan int) {
	x, y := 0, 1
	for {
		select {
		case c <- x:
			x, y = y, x+y
		case <-quit:
			fmt.Println("quit")
			return
		}
	}
}

//默认选择
func Test_default_selection(t *testing.T) {
	println("当 select 中的其它分支都没有准备好时，default 分支就会执行。")
	println("为了在尝试发送或者接收时不发生阻塞，可使用 default 分支：")
	tick := time.Tick(100 * time.Millisecond)
	boom := time.After(500 * time.Millisecond)
	for {
		select {
		case <-tick:
			fmt.Println("tick.")
		case <-boom:
			fmt.Println("BOOM!")
			return
		default:
			fmt.Println("    .")
			time.Sleep(50 * time.Millisecond)
		}
	}

}

//练习：等价二叉查找树
// func Test_exercise_equivalent_binary_trees(t *testing.T) {
// 	ch := make(chan int)
// 	t1 := tree.New(1)
// 	t2 := tree.New(2)

// 	//从信道中打印10个值
// 	go Walk(t1, ch)
// 	for i := 0; i < 10; i++ {
// 		fmt.Println(<-ch)
// 	}

// 	//对tree1和tree1进行比较
// 	fmt.Println("tree 1 == tree 1:", Same(t1, t2))
// 	fmt.Println("tree 1 == tree 2:", Same(t1, t2))
// }

// //Walk 步进 tree t 将所有的值从 tree 发送到 channel ch。
// func Walk(t *tree.Tree, ch chan int) {
// 	if t.Left != nil {
// 		Walk(t.Left, ch)
// 	}
// 	ch <- t.Value
// 	if t.Right != nil {
// 		Walk(t.Right, ch)
// 	}
// 	return
// }

// // Same 检测树 t1 和 t2 是否含有相同的值。
// func Same(t1, t2 *tree.Tree) bool {
// 	ch1, ch2 := make(chan int), make(chan int)
// 	go Walk(t1, ch1)
// 	go Walk(t2, ch2)
// 	if <-ch1 == <-ch2 {
// 		return true
// 	}
// 	return false
// }

//sync.Mutex
func Test_mutex_counter(t *testing.T) {
	println("这里涉及的概念叫做 *互斥（mutual*exclusion）* ，我们通常使用 *互斥锁（Mutex）* 这一数据结构来提供这种机制。")
	println("Go 标准库中提供了 sync.Mutex 互斥锁类型及其两个方法： Lock Unlock")
	println("我们可以通过在代码前调用 Lock 方法，在代码后调用 Unlock 方法来保证一段代码的互斥执行")
	println("我们也可以用 defer 语句来保证互斥锁一定会被解锁")

	c := SafeCounter{v: make(map[string]int)}
	for i := 0; i < 1000; i++ {
		go c.Inc("somekey")
	}

	time.Sleep(time.Second)
	fmt.Println(c.Value("somekey"))
}

type SafeCounter struct {
	//SafeCounter 的并发使用是安全的
	v   map[string]int
	mux sync.Mutex
}

func (c *SafeCounter) Inc(key string) {
	//Inc 增加给定key的计数器的值
	c.mux.Lock()
	//Lock 之后同一时刻只有一个goroutine能访问c.V
	c.v[key]++
	c.mux.Unlock()
}

func (c *SafeCounter) Value(key string) int {
	//Value返回给定的key 的计数器的当前值
	c.mux.Lock()
	//Lock 之后同一时刻只有一个goroutine能访问c.V
	defer c.mux.Unlock()
	return c.v[key]
}

//练习：Web 爬虫
func Test_exercise_web_crawler(t *testing.T) {
	println("修改 Crawl 函数来并行地抓取 URL，并且保证不重复。")
	crawlerMap := CrawlerMap{v: make(map[string]string)}
	go Crawl("https://golang.org/", 4, fetcher, crawlerMap)
	time.Sleep(time.Second)
	for url, body := range crawlerMap.v {
		fmt.Printf("found: %s %q\n", url, body)
	}
}

type Fetcher interface {
	// Fetch 返回 URL 的 body 内容，并且将在这个页面上找到的 URL 放到一个 slice 中。
	Fetch(url string) (body string, urls []string, err error)
}

// Crawl 使用 fetcher 从某个 URL 开始递归的爬取页面，直到达到最大深度。
func Crawl(url string, depth int, fetcher Fetcher, crawlerMap CrawlerMap) {
	if depth <= 0 {
		return
	}
	body, urls, err := fetcher.Fetch(url)
	if err != nil {
		crawlerMap.put(url, err.Error())
		return
	}
	crawlerMap.put(url, body)
	for _, u := range urls {
		if _, ok := crawlerMap.v[u]; !ok {
			go Crawl(u, depth-1, fetcher, crawlerMap)
		}
	}
	return
}

// fakeFetcher 是返回若干结果的 Fetcher。
type fakeFetcher map[string]*fakeResult

type fakeResult struct {
	body string
	urls []string
}

func (f fakeFetcher) Fetch(url string) (string, []string, error) {
	if res, ok := f[url]; ok {
		return res.body, res.urls, nil
	}
	return "", nil, fmt.Errorf("not found: %s", url)
}

var fetcher = fakeFetcher{
	"https://golang.org/": &fakeResult{
		"The Go Programming Language",
		[]string{
			"https://golang.org/pkg/",
			"https://golang.org/cmd/",
		},
	},
	"https://golang.org/pkg/": &fakeResult{
		"Packages",
		[]string{
			"https://golang.org/",
			"https://golang.org/cmd/",
			"https://golang.org/pkg/fmt/",
			"https://golang.org/pkg/os/",
		},
	},
	"https://golang.org/pkg/fmt/": &fakeResult{
		"Package fmt",
		[]string{
			"https://golang.org/",
			"https://golang.org/pkg/",
		},
	},
	"https://golang.org/pkg/os/": &fakeResult{
		"Package os",
		[]string{
			"https://golang.org/",
			"https://golang.org/pkg/",
		},
	},
}

type CrawlerMap struct {
	v   map[string]string
	mux sync.Mutex
}

func (c *CrawlerMap) put(key, value string) {
	c.mux.Lock()
	c.v[key] = value
	c.mux.Unlock()
}

func (c *CrawlerMap) get(key string) string {
	c.mux.Lock()
	value := c.v[key]
	c.mux.Unlock()
	return value
}
