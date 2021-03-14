package leetcode

import (
	"testing"
)

/**
 * 705. 设计哈希集合
 */

func Test_leetcode_705(t *testing.T) {
	myHashSet := Constructor705()
	myHashSet.Add(1)      // set = [1]
	myHashSet.Add(2)      // set = [1, 2]
	myHashSet.Contains(1) // 返回 True
	myHashSet.Contains(3) // 返回 False ，（未找到）
	myHashSet.Add(2)      // set = [1, 2]
	myHashSet.Contains(2) // 返回 True
	myHashSet.Remove(2)   // set = [1]
	myHashSet.Contains(2) // 返回 False ，（已移除）

}

type MyHashSet struct {
	bitset []uint64
}

/** Initialize your data structure here. */
func Constructor705() MyHashSet {
	return MyHashSet{bitset: []uint64{}}
}

func (this *MyHashSet) Add(key int) {
	bit := key % 64
	length := key / 64
	for i := len(this.bitset); i <= length; i++ {
		this.bitset = append(this.bitset, 0)
	}
	this.bitset[length] = this.bitset[length] | (1 << bit)
}

func (this *MyHashSet) Remove(key int) {
	bit := key % 64
	length := key / 64
	if length >= len(this.bitset) {
		return
	}
	this.bitset[length] = this.bitset[length] & ^(1 << bit)
}

/** Returns true if this set contains the specified element */
func (this *MyHashSet) Contains(key int) bool {
	bit := key % 64
	length := key / 64
	if length >= len(this.bitset) {
		return false
	}
	return this.bitset[length]&(1<<bit) != 0
}
