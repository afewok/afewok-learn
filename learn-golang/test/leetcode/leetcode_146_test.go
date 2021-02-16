package leetcode

import (
	"fmt"
	"testing"
)

/**
 * 146. LRU 缓存机制
 */

func Test_leetcode_146(t *testing.T) {
	lRUCache := Constructor146(2)
	fmt.Println(lRUCache.Get146(2))
	lRUCache.Put146(2, 6)
	fmt.Println(lRUCache.Get146(1))
	lRUCache.Put146(1, 5)
	lRUCache.Put146(1, 2)
	fmt.Println(lRUCache.Get146(1))
	fmt.Println(lRUCache.Get146(2))
}

type LRUCache struct {
	Key      int
	Val      int
	Next     *LRUCache
	Capacity int
}

func Constructor146(capacity int) LRUCache {
	return LRUCache{-1, -1, nil, capacity}
}

func (this *LRUCache) Get146(key int) int {
	p, q := this.Next, this
	for p != nil {
		if p.Key == key {
			if this.Next == p {
				return this.Next.Val
			}
			q.Next, p.Next, this.Next = p.Next, this.Next, p
			p.Capacity, p, q.Capacity = this.Capacity, p.Next, q.Capacity-1
			for p != q {
				p.Capacity--
				p = p.Next
			}
			return this.Next.Val
		}
		q, p = p, p.Next

	}
	return -1
}

func (this *LRUCache) Put146(key int, value int) {
	this.Next = &LRUCache{key, value, this.Next, this.Capacity}
	p := this.Next
	for p.Next != nil {
		if p.Next.Key == key {
			p.Next = p.Next.Next
			break
		}
		p.Next.Capacity--
		if p.Next.Capacity == 0 {
			p.Next = nil
			break
		}
		p = p.Next
	}
}
