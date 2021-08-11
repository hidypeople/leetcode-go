package tasks

import ll "leetcode/linkedList"

// Merge list of ordered linked lists into one
func MergeKLists(lists []*ll.ListNode) *ll.ListNode {
	n := len(lists)
	var result, curr *ll.ListNode = nil, nil
	currValue := 100000
	currents := make([]*ll.ListNode, n)
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
func fastRemove(nodes []*ll.ListNode, i int) []*ll.ListNode {
	nodes[i] = nodes[len(nodes)-1]
	return nodes[:len(nodes)-1]
}

// Find minimum value withing given nodes (currVal used as early exit optimization)
func findMinCurrent(nodes []*ll.ListNode, currVal int) (*ll.ListNode, int) {
	var minNode *ll.ListNode = nil
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
