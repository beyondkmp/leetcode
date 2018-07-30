// This example demonstrates a priority queue built using the heap interface.
package main

import (
	"container/heap"
	"fmt"
)

// An Item is something we manage in a priority queue.
type Item struct {
	value    string // The value of the item; arbitrary.
	priority int    // The priority of the item in the queue.
	// The index is needed by update and is maintained by the heap.Interface methods.
	index int // The index of the item in the heap.
}

// A PriorityQueue implements heap.Interface and holds Items.
type PriorityQueue []*Item

func (pq PriorityQueue) Len() int { return len(pq) }

func (pq PriorityQueue) Less(i, j int) bool {
	// We want Pop to give us the highest, not lowest, priority so we use greater than here.
	if pq[i].priority > pq[j].priority {
		return true
	}
	if pq[i].priority == pq[j].priority && pq[i].value < pq[j].value {
		return true
	}
	return false
}

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].index = i
	pq[j].index = j

}

func (pq *PriorityQueue) Push(x interface{}) {
	n := len(*pq)
	item := x.(*Item)
	item.index = n
	*pq = append(*pq, item)
}

func (pq *PriorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	item.index = -1 // for safety
	*pq = old[0 : n-1]
	return item
}

func topKFrequent(words []string, k int) []string {
	freq := make(map[string]int)
	for _, w := range words {
		freq[w]++
	}
	pq := make(PriorityQueue, len(freq))
	i := 0
	for value, priority := range freq {
		pq[i] = &Item{
			value:    value,
			priority: priority,
			index:    i,
		}
		i++
	}
	heap.Init(&pq)

	var result []string
	i = 0
	for i < k {
		item := heap.Pop(&pq).(*Item)
		result = append(result, item.value)
		i++
	}
	return result
}

// This example creates a PriorityQueue with some items, adds and manipulates an item,
// and then removes the items in priority order.
func main() {
	words := []string{"the", "day", "is", "sunny", "the", "the", "the", "sunny", "is", "is"}
	k := 4
	fmt.Println(topKFrequent(words, k))

	words = []string{"i", "love", "leetcode", "i", "love", "coding"}
	k = 3
	fmt.Println(topKFrequent(words, k))
}
