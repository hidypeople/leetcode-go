package task0023

import . "leetcode/linkedList"

// You are given an array of k linked-lists lists, each linked-list is sorted in ascending order.
// Merge all the linked-lists into one sorted linked-list and return it.
//
// Constraints:
//   k == lists.length
//   0 <= k <= 10^4
//   0 <= lists[i].length <= 500
//   -10^4 <= lists[i][j] <= 10^4
//   lists[i] is sorted in ascending order.
//   The sum of lists[i].length won't exceed 10^4.
//
// Results:
//   Runtime: 32 ms, faster than 36.97% of Go online submissions for Merge k Sorted Lists.
//   Memory Usage: 5.8 MB, less than 40.72% of Go online submissions for Merge k Sorted Lists.
func mergeKLists(lists []*ListNode) *ListNode {
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
