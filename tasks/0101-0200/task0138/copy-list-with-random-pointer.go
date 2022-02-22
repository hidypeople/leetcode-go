package task0138

type Node struct {
	Val    int
	Next   *Node
	Random *Node
}

// A linked list of length n is given such that each node contains an additional random pointer,
// which could point to any node in the list, or null.
// Construct a deep copy of the list. The deep copy should consist of exactly n brand new nodes,
// where each new node has its value set to the value of its corresponding original node.
// Both the next and random pointer of the new nodes should point to new nodes in the copied
// list such that the pointers in the original list and copied list represent the same list state.
// None of the pointers in the new list should point to nodes in the original list.
// For example, if there are two nodes X and Y in the original list, where X.random --> Y,
// then for the corresponding two nodes x and y in the copied list, x.random --> y.
// Return the head of the copied linked list.
// The linked list is represented in the input/output as a list of n nodes.
// Each node is represented as a pair of [val, random_index] where:
// val: an integer representing Node.val
// random_index: the index of the node (range from 0 to n-1) that the random pointer
// points to, or null if it does not point to any node.
// Your code will only be given the head of the original linked list.
//
// Constraints:
//   0 <= n <= 1000
//   -10000 <= Node.val <= 10000
//   Node.random is null or is pointing to some node in the linked list.
//
// Results:
//   Runtime: 0 ms, faster than 100.00% of Go online submissions for Copy List with Random Pointer.
//   Memory Usage: 3.7 MB, less than 44.60% of Go online submissions for Copy List with Random Pointer.
func copyRandomList(head *Node) *Node {
	if head == nil {
		return nil
	}

	nodeToIndexOld := map[*Node]int{head: 0} // oldNode -> idx
	indexToNodeNew := map[int]*Node{}        // idx -> newNode

	// Build a copy without random part
	var headNew *Node = nil
	var currNew *Node = nil
	curr := head
	i := 0
	for curr != nil {
		new := &Node{Val: curr.Val}
		if headNew == nil {
			headNew = new
		} else {
			currNew.Next = new
		}

		nodeToIndexOld[curr] = i
		indexToNodeNew[i] = new
		currNew = new
		curr = curr.Next
		i++
	}

	// Attach random part to each node
	curr, currNew = head, headNew
	for curr != nil {
		if curr.Random != nil {
			currNew.Random = indexToNodeNew[nodeToIndexOld[curr.Random]]
		}
		currNew = currNew.Next
		curr = curr.Next
	}

	return headNew
}
