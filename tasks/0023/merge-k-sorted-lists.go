package tasks

import . "leetcode/linkedList"

// Merge list of ordered linked lists into one ordered list
func MergeKLists(lists []*ListNode) *ListNode {
	n := len(lists)
	var result, curr *ListNode = nil, nil
	currValue := 100000
	currents := make([]*ListNode, n)
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
func fastRemove(nodes []*ListNode, i int) []*ListNode {
	nodes[i] = nodes[len(nodes)-1]
	return nodes[:len(nodes)-1]
}

// Find minimum value withing given nodes (currVal used as early exit optimization)
func findMinCurrent(nodes []*ListNode, currVal int) (*ListNode, int) {
	var minNode *ListNode = nil
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
