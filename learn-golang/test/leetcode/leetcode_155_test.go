package leetcode

import (
	"testing"
)

/**
 * 155. 最小栈
 *
 */
func Test_leetcode_155(t *testing.T) {
	minStack := Constructor155()
	minStack.Push(-2)
	minStack.Push(0)
	minStack.Push(-3)
	minStack.GetMin() // 返回 -3.
	minStack.Pop()
	minStack.Top()    // 返回 0.
	minStack.GetMin() // 返回 -2.
}

type MinStack struct {
	min []int
	arr []int
}

/** initialize your data structure here. */
func Constructor155() MinStack {
	return MinStack{make([]int, 0), make([]int, 0)}
}

func (this *MinStack) Push(x int) {
	this.arr = append(this.arr, x)
	for i := 0; i < len(this.min); i++ {
		if x < this.min[i] {
			this.min = append(this.min[:i], append([]int{x}, this.min[i:]...)...)
			return
		}
	}
	this.min = append(this.min, x)
}

func (this *MinStack) Pop() {
	x := this.arr[len(this.arr)-1]
	this.arr = this.arr[:len(this.arr)-1]
	for i := 0; i < len(this.min); i++ {
		if x == this.min[i] {
			this.min = append(this.min[:i], this.min[i+1:]...)
		}
	}
}

func (this *MinStack) Top() int {
	return this.arr[len(this.arr)-1]
}

func (this *MinStack) GetMin() int {
	return this.min[0]
}
