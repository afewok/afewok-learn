package leetcode

import (
	"testing"
)

/**
 * 706. 设计哈希映射
 *
 */
func Test_leetcode_706(t *testing.T) {
	myHashMap := Constructor706()
	myHashMap.Put(1, 1) // myHashMap 现在为 [[1,1]]
	myHashMap.Put(2, 2) // myHashMap 现在为 [[1,1], [2,2]]
	myHashMap.Get(1)    // 返回 1 ，myHashMap 现在为 [[1,1], [2,2]]
	myHashMap.Get(3)    // 返回 -1（未找到），myHashMap 现在为 [[1,1], [2,2]]
	myHashMap.Put(2, 1) // myHashMap 现在为 [[1,1], [2,1]]（更新已有的值）
	myHashMap.Get(2)    // 返回 1 ，myHashMap 现在为 [[1,1], [2,1]]
	myHashMap.Remove(2) // 删除键为 2 的数据，myHashMap 现在为 [[1,1]]
	myHashMap.Get(2)    // 返回 -1（未找到），myHashMap 现在为 [[1,1]]
}

type HashKV struct {
	Key   int
	Value int
	Next  *HashKV
}
type MyHashMap struct {
	data []*HashKV
}

/** Initialize your data structure here. */
func Constructor706() MyHashMap {
	data := make([]*HashKV, 32)
	for i := 0; i < 32; i++ {
		data[i] = &HashKV{-1, -1, nil}
	}
	return MyHashMap{data}
}

/** value will always be non-negative. */
func (this *MyHashMap) Put(key int, value int) {
	kv := this.data[key%32]
	for kv.Next != nil {
		if kv.Next.Key == key {
			kv.Next.Value = value
			return
		}
		kv = kv.Next
	}
	kv.Next = &HashKV{key, value, nil}
}

/** Returns the value to which the specified key is mapped, or -1 if this map contains no mapping for the key */
func (this *MyHashMap) Get(key int) int {
	kv := this.data[key%32]
	for kv.Next != nil {
		if kv.Next.Key == key {
			return kv.Next.Value
		}
		kv = kv.Next
	}
	return -1
}

/** Removes the mapping of the specified value key if this map contains a mapping for the key */
func (this *MyHashMap) Remove(key int) {
	kv := this.data[key%32]
	for kv.Next != nil {
		if kv.Next.Key == key {
			kv.Next = kv.Next.Next
			return
		}
		kv = kv.Next
	}
}
