package leetcode

import (
	"fmt"
	"testing"
)

/**
 * 225. 用队列实现栈
 */

func Test_leetcode_225(t *testing.T) {
	obj := Constructor()
	obj.PushStack(99)
	fmt.Println(obj.TopStack(), obj.PopStack(), obj.EmptyStack())
}

type MyStack struct {
	Queue []int
}

/** Initialize your data structure here. */
func Constructor() MyStack {
	return MyStack{}
}

/** Push element x onto stack. */
func (stack *MyStack) PushStack(x int) {
	stack.Queue = append(stack.Queue, x)
}

/** Removes the element on top of the stack and returns that element. */
func (stack *MyStack) PopStack() int {
	result := stack.Queue[len(stack.Queue)-1]
	stack.Queue = stack.Queue[:len(stack.Queue)-1]
	return result
}

/** Get the top element. */
func (stack *MyStack) TopStack() int {
	return stack.Queue[len(stack.Queue)-1]
}

/** Returns whether the stack is empty. */
func (stack *MyStack) EmptyStack() bool {
	if len(stack.Queue) == 0 {
		return true
	}
	return false
}
