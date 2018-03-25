package main

import (
	"fmt"
)

/**
 * Your LRUCache object will be instantiated and called as such:
 * obj := Constructor(mxCapacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */
type LRUCache struct {
	mxCapacity int
	lruMap     map[int]*DoubleLinked
	pqFirst    *DoubleLinked
	pqLast     *DoubleLinked
}

type CacheEntry struct {
	key   int
	value int
}
type DoubleLinked struct {
	forward  *DoubleLinked
	backward *DoubleLinked
	entry    CacheEntry
}

func Constructor(capacity int) LRUCache {
	lruMap := make(map[int]*DoubleLinked)
	return LRUCache{lruMap: lruMap, mxCapacity: capacity}

}

func (this *LRUCache) Get(key int) int {
	dlEntry, ok := this.lruMap[key]
	if !ok {
		return -1
	}
	tempF := dlEntry.forward
	if tempF != nil { //equal to nil for forward pointer - already at the start of the queue
		if this.pqLast == dlEntry {
			this.pqLast = tempF
		}
		tempB := dlEntry.backward
		if tempB != nil {
			tempF.backward, tempB.forward = tempB, tempF
		} else {
			tempF.backward = nil
		}
		dlEntry.forward = nil
		dlEntry.backward = this.pqFirst
		this.pqFirst.forward = dlEntry
		this.pqFirst = dlEntry
	}
	return dlEntry.entry.value
}

func (this *LRUCache) Put(key int, value int) {
	dlNewEntry, ok := this.lruMap[key]
	if ok {
		dlNewEntry.entry.value = value
		this.Get(key)
		return
	}
	if len(this.lruMap) == this.mxCapacity {
		delete(this.lruMap, this.pqLast.entry.key)
		if this.pqLast != nil {
			this.pqLast = this.pqLast.forward
			this.pqLast.backward = nil
		}
	}
	dlNewEntry = &DoubleLinked{
		entry: CacheEntry{
			value: value,
			key:   key,
		},
	}
	this.lruMap[key] = dlNewEntry

	if this.pqFirst == nil {
		this.pqFirst = dlNewEntry
		this.pqLast = dlNewEntry
	} else {
		this.pqFirst.forward = dlNewEntry
		dlNewEntry.backward = this.pqFirst
		this.pqFirst = dlNewEntry
	}
}

/* // A PriorityQueue implements heap.Interface and holds Items.
type PriorityQueue []*CacheEntry

func (pq PriorityQueue) Len() int { return len(pq) }

func (pq PriorityQueue) Less(i, j int) bool {
	// We want Pop to give us the highest, not lowest, priority so we use greater than here.
	return pq[i].priority <= pq[j].priority
}

//func (pq PriorityQueue) Remove (x interface{}){
//	heap.Fix()
//
//}
func (pq *PriorityQueue) Push(x interface{}) {
	item := x.(*CacheEntry)
	item.index = 0
	*pq = append(*pq, item)
}

func (pq *PriorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[0]
	item.index = -1 // for safety
	*pq = old[1: n]
	return item
}

// update modifies the priority and value of an Item in the queue.
func (pq *PriorityQueue) update(entry *CacheEntry, value int, priority int) {
	entry.value = value
	entry.priority = priority
	heap.Fix(pq, entry.index)
}
func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].index = i
	pq[j].index = j
}
*/
func main() {
	//	["LRUCache","put","put","put","put","get","get"]
	//[[2],[2,1],[1,1],[2,3],[4,1],[1],[2]]
	obj := Constructor(2)
	obj.Put(2, 1)
	obj.Put(1, 1)
	obj.Put(2, 3)
	obj.Put(4, 1)
	check(-1, obj.Get(1))

	check(3, obj.Get(2))
	//obj := Constructor(2)
	//obj.Put(1, 1)
	//obj.Put(2, 2)
	//check(1, obj.Get(1))  // returns 1
	//obj.Put(3, 3)         // evicts key 2
	//check(-1, obj.Get(2)) // returns -1 (not found)
	//obj.Put(4, 4)         // evicts key 1
	//check(-1, obj.Get(1)) // returns -1 (not found)
	//check(3, obj.Get(3))  // returns 3
	//check(4, obj.Get(4))  // returns 4
}

func check(should int, val int) {
	fmt.Printf("should be %d but it is %d\n", should, val)
}

func (c CacheEntry) String() string {
	return fmt.Sprintf("key : %d value %d", c.key, c.value)
}

func (d DoubleLinked) String() string {
	return fmt.Sprintf("key : %d value %d", d.entry.key, d.entry.value)
}
