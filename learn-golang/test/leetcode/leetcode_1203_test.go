package leetcode

import (
	"container/heap"
	"fmt"
	"testing"
)

/**
 * 1203. 项目管理
 */

func Test_leetcode_1203(t *testing.T) {
	fmt.Println(sortItems(8, 2, []int{-1, -1, 1, 0, 0, 1, 0, -1}, [][]int{{}, {6}, {5}, {6}, {3, 6}, {}, {}, {}}))
	fmt.Println(sortItems(8, 2, []int{-1, -1, 1, 0, 0, 1, 0, -1}, [][]int{{}, {6}, {5}, {6}, {3}, {}, {4}, {}}))
}

/*拓扑排序加优先队列*/
func sortItems(n int, m int, group []int, beforeItems [][]int) []int {
	defer timeCost()()
	maxGroup := m
	for i, v := range group {
		if v == -1 {
			group[i] = maxGroup
			maxGroup++
		}
	}
	groupItems := make([][]int, maxGroup)
	for i, v := range group {
		groupItems[v] = append(groupItems[v], i)
	}
	//项目顶点入度及边
	itemInDegree := make([]int, n)
	itemEdges := make(map[int][]int)
	//小组顶点入度及边
	groupInDegree := make([]int, maxGroup)
	groupEdges := make(map[int]map[int]bool)
	//初始化 1.项目顶点入度及边 2.小组顶点入度及边
	for i := 0; i < len(beforeItems); i++ {
		for _, v := range beforeItems[i] {
			if _, ok := itemEdges[v]; !ok {
				itemEdges[v] = make([]int, 0)
			}
			itemEdges[v] = append(itemEdges[v], i)
			itemInDegree[i]++
			if group[v] != group[i] {
				if _, ok := groupEdges[group[v]][group[i]]; !ok {
					if _, ok := groupEdges[group[v]]; !ok {
						groupEdges[group[v]] = make(map[int]bool)
					}
					groupEdges[group[v]][group[i]] = true
					groupInDegree[group[i]]++
				}
			}
		}
	}
	/*
	 *小组拓扑排序，前置分组排在前面，优先队列， 入度都是0的时候，优先级使用小组编号优先输出
	 */
	ans := []int{}
	pq := &PriorityQueue{}
	for i, v := range groupInDegree {
		if v == 0 {
			item := &Item{
				value:    i,
				priority: i,
			}
			heap.Push(pq, item)
		}
	}
	for pq.Len() > 0 {
		item := heap.Pop(pq).(*Item)
		v := item.value
		ans = append(ans, v)
		for w := range groupEdges[v] {
			groupInDegree[w]--
			if groupInDegree[w] == 0 {
				item := &Item{
					value:    w,
					priority: w,
				}
				heap.Push(pq, item)
			}
		}
	}
	if len(ans) < maxGroup {
		return []int{}
	}
	//对项目拓扑排序，优先队列， 优先级使用小组编号，确保同一小组的项目彼此相邻的输出
	res := []int{}
	for i := 0; i < len(ans); i++ {
		for _, w := range groupItems[ans[i]] {
			if itemInDegree[w] == 0 {
				item := &Item{
					value:    w,
					priority: group[w],
				}
				heap.Push(pq, item)
			}
		}
		for pq.Len() > 0 {
			item := heap.Pop(pq).(*Item)
			v := item.value
			itemInDegree[v]--
			res = append(res, v)
			for _, w := range itemEdges[v] {
				itemInDegree[w]--
				if itemInDegree[w] == 0 && group[v] == group[w] {
					item := &Item{
						value:    w,
						priority: group[w],
					}
					heap.Push(pq, item)
				}
			}
		}
	}
	if len(res) < n {
		return []int{}
	}
	return res
}

/*堆 优先队列*/
type Item struct {
	value    int // The value of the item; arbitrary.
	priority int // The priority of the item in the queue.
}

type PriorityQueue []*Item

func (pq PriorityQueue) Len() int { return len(pq) }

func (pq PriorityQueue) Less(i, j int) bool {
	return pq[i].priority < pq[j].priority
}

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
}

func (pq *PriorityQueue) Push(x interface{}) {
	item := x.(*Item)
	*pq = append(*pq, item)
}

func (pq *PriorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	old[n-1] = nil // avoid memory leak
	*pq = old[0 : n-1]
	return item
}
