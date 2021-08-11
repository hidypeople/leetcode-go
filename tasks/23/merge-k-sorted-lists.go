package tasks

import ll "leetcode/linkedList"

// Merge list of ordered linked lists into one
func MergeKLists(lists []*ll.ListNodeInt) *ll.ListNodeInt {
	n := len(lists)
	var result, curr *ll.ListNodeInt = nil, nil
	currValue := 100000
	currents := make([]*ll.ListNodeInt, n)
	copy(currents, lists)
	for len(currents) > 0 {
		next, nextIndex := findMinCurrent(currents, currValue)
		currValue = next.Val
		currents[nextIndex] = currents[nextIndex].Next
		if currents[nextIndex] == nil {
			currents = fastRemove(currents, nextIndex)
		}
		if result == nil {
			result, curr = next, next
		} else {
			curr.Next = next
			curr = next
		}
	}
	return result
}

// Fast remove element without keeping the order
func fastRemove(nodes []*ll.ListNodeInt, i int) []*ll.ListNodeInt {
	nodes[i] = nodes[len(nodes)-1]
	return nodes[:len(nodes)-1]
}

// Find minimum value withing given nodes (currVal used as early exit optimization)
func findMinCurrent(nodes []*ll.ListNodeInt, currVal int) (*ll.ListNodeInt, int) {
	var minNode *ll.ListNodeInt = nil
	var minIndex int = -1
	for i := 0; i < len(nodes); i++ {
		if i == 0 || nodes[i].Val < minNode.Val {
			minNode = nodes[i]
			minIndex = i
			// Fast return
			if minNode.Val == currVal {
				return minNode, minIndex
			}
		}
	}
	return minNode, minIndex
}
