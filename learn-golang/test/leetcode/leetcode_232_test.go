package leetcode

import (
	"fmt"
	"testing"
)

/**
 * 232. 用栈实现队列
 */

func Test_leetcode_232(t *testing.T) {
	obj := ConstructorQueue()
	obj.PushQueue(1)
	obj.PushQueue(2)
	fmt.Println(obj.PeekQueue(), obj.PopQueue(), obj.EmptyQueue())
}

type MyQueue struct {
	Stack []int
}

/** Initialize your data structure here. */
func ConstructorQueue() MyQueue {
	return MyQueue{}
}

/** Push element x to the back of queue. */
func (queue *MyQueue) PushQueue(x int) {
	queue.Stack = append(queue.Stack, x)

}

/** Removes the element from in front of queue and returns that element. */
func (queue *MyQueue) PopQueue() int {
	x := queue.Stack[0]
	queue.Stack = queue.Stack[1:]
	return x
}

/** Get the front element. */
func (queue *MyQueue) PeekQueue() int {
	return queue.Stack[0]
}

/** Returns whether the queue is empty. */
func (queue *MyQueue) EmptyQueue() bool {
	if len(queue.Stack) == 0 {
		return true
	}
	return false
}
