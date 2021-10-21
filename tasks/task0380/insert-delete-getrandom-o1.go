package task0380

import (
	"math/rand"
)

// Implement the RandomizedSet class:
//   RandomizedSet() Initializes the RandomizedSet object.
//   bool insert(int val)
//     Inserts an item val into the set if not present.
//     Returns true if the item was not present, false otherwise.
//   bool remove(int val)
//     Removes an item val from the set if present.
//     Returns true if the item was present, false otherwise.
//   int getRandom()
//     Returns a random element from the current set of elements
//     (it's guaranteed that at least one element exists when this
//     method is called). Each element must have the same probability of being returned.
// You must implement the functions of the class such that each function works in average O(1) time complexity.
//
// Constraints:
//   -2^31 <= val <= 2^31 - 1
//   At most 2 * 10^5 calls will be made to insert, remove, and getRandom.
//   There will be at least one element in the data structure when getRandom is called.
//
// Results:
//   Runtime: 156 ms, faster than 99.44% of Go online submissions for Insert Delete GetRandom O(1). *Unstable*
//   Memory Usage: 57.1 MB, less than 15.64% of Go online submissions for Insert Delete GetRandom O(1).
type RandomizedSet struct {
	len    int
	keys   []int
	values map[int]int // key: the value, value: the index in this.keys for that value
}

func Constructor() RandomizedSet {
	return RandomizedSet{
		len:    0,
		keys:   []int{},
		values: map[int]int{},
	}
}

func (rs *RandomizedSet) Insert(val int) bool {
	if _, ok := rs.values[val]; ok {
		return false
	}
	rs.keys = append(rs.keys, val)
	rs.values[val] = len(rs.keys) - 1
	rs.len++
	return true
}

func (rs *RandomizedSet) Remove(val int) bool {
	currentIdx, ok := rs.values[val]
	if !ok {
		return false
	}
	lastVal := rs.keys[rs.len-1]
	if lastVal != val {
		rs.keys[currentIdx] = lastVal
		rs.keys = rs.keys[:rs.len-1]
		rs.values[lastVal] = currentIdx
	} else {
		rs.keys = rs.keys[:rs.len-1]
	}
	delete(rs.values, val)
	rs.len--
	return true
}

func (rs *RandomizedSet) GetRandom() int {
	return rs.keys[rand.Intn(rs.len)]
}
