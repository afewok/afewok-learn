package learnblog

import (
	"fmt"
	"testing"
)

func Test_Panic_Recover(t *testing.T) {
	recoverTest()
	fmt.Println("Returned normally from recoverTest.")
}

func recoverTest() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in recoverTest", r)
		}
	}()
	fmt.Println("Calling panicTest.")
	panicTest(0)
	fmt.Println("Returned normally from panicTest.")
}

func panicTest(i int) {
	if i > 3 {
		fmt.Println("Panicking!")
		panic(fmt.Sprintf("%v", i))
	}
	defer fmt.Println("Defer in panicTest", i)
	fmt.Println("Printing in panicTest", i)
	panicTest(i + 1)
}
