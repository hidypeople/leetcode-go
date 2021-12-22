package task0146

// Design a data structure that follows the constraints of a Least Recently Used (LRU) cache.
// Implement the LRUCache class:
//   LRUCache(int capacity) Initialize the LRU cache with positive size capacity.
//   int get(int key) Return the value of the key if the key exists, otherwise return -1.
//   void put(int key, int value) Update the value of the key if the key exists.
//      Otherwise, add the key-value pair to the cache. If the number of keys exceeds the capacity
//      from this operation, evict the least recently used key.
//   The functions get and put must each run in O(1) average time complexity.
//
// Constraints:
//   1 <= capacity <= 3000
//   0 <= key <= 10^4
//   0 <= value <= 10^5
//   At most 2 * 10^5 calls will be made to get and put.
//
// Results:
//   Runtime: 536 ms, faster than 98.84% of Go online submissions for LRU Cache.
//   Memory Usage: 85.9 MB, less than 28.93% of Go online submissions for LRU Cache.
type ListNode struct {
	Key  int
	Val  int
	Prev *ListNode
	Next *ListNode
}

type LRUCache struct {
	Capacity int
	Length   int
	Head     *ListNode
	Tail     *ListNode
	Nodes    map[int]*ListNode
}

func Constructor(capacity int) LRUCache {
	return LRUCache{
		Capacity: capacity,
		Length:   0,
		Head:     nil,
		Tail:     nil,
		Nodes:    make(map[int]*ListNode, capacity),
	}
}

func (cache *LRUCache) RaiseKey(key int) {
	node := cache.Nodes[key]
	if node == cache.Head {
		return
	}
	// Reconnect prev - node - next => prev - next
	prev, next := node.Prev, node.Next
	prev.Next = next
	if next == nil {
		// If node is the last node => update cache.Tail
		cache.Tail = prev
	} else {
		next.Prev = prev
	}

	// Move node to cache.Head
	cache.Head.Prev = node
	node.Next = cache.Head
	node.Prev = nil
	cache.Head = node
}

func (cache *LRUCache) EvictKey() {
	delete(cache.Nodes, cache.Tail.Key)
	if cache.Capacity == 1 {
		cache.Length = 0
		cache.Tail = nil
		cache.Head = nil
	} else {
		cache.Tail = cache.Tail.Prev
		cache.Tail.Next = nil
		cache.Length--
	}

}

func (cache *LRUCache) Get(key int) int {
	node, ok := cache.Nodes[key]
	if !ok {
		return -1
	}
	if cache.Length > 1 {
		cache.RaiseKey(key)
	}
	return node.Val
}

func (cache *LRUCache) Put(key int, value int) {
	node, ok := cache.Nodes[key]
	if ok {
		// Existing key: update it
		if cache.Length > 1 {
			cache.RaiseKey(key)
		}
		node.Val = value
		return
	}

	if cache.Length == cache.Capacity {
		// Cache is full => evict the less recent item
		cache.EvictKey()
	}

	// Add node to the tail
	newNode := &ListNode{
		Key:  key,
		Val:  value,
		Prev: nil,
		Next: nil,
	}
	cache.Nodes[key] = newNode
	if cache.Length == 0 {
		// cache empty => init with 1 element
		cache.Head = newNode
		cache.Tail = newNode
	} else {
		// cache not empty => raise new node
		head := cache.Head
		cache.Head = newNode
		head.Prev, newNode.Next = newNode, head
	}
	cache.Length++
}
