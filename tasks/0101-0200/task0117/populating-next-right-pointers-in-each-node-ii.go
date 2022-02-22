package task0117

// Given a binary tree
// struct Node {
// 	int val;
// 	Node *left;
// 	Node *right;
// 	Node *next;
// }
// Populate each next pointer to point to its next right node. If there is no next right node,
// the next pointer should be set to NULL.
// Initially, all next pointers are set to NULL.
//
// Constraints:
//   The number of nodes in the tree is in the range [0, 6000].
//   -100 <= Node.val <= 100
//
// Results:
//   Runtime: 0 ms, faster than 100.00% of Go online submissions for Populating Next Right Pointers in Each Node II.
//   Memory Usage: 6.3 MB, less than 100.00% of Go online submissions for Populating Next Right Pointers in Each Node II.
type Node struct {
	Val   int
	Left  *Node
	Right *Node
	Next  *Node
}

func connect(root *Node) *Node {
	if root == nil {
		return nil
	}
	connectCurrent(root)
	return root
}

func attachNext(parent *Node, node *Node) {
	// Use the fact that parent level is completed and we can navigate through the same level parents using Next property
	p := parent
	for p != nil {
		if p.Left != nil && p.Left != node && p.Right != node {
			node.Next = p.Left
			return
		}
		if p.Right != nil && p.Right != node {
			node.Next = p.Right
			return
		}
		p = p.Next
	}
}

func connectCurrent(node *Node) {
	if node.Right != nil {
		attachNext(node, node.Right)
	}
	if node.Left != nil {
		attachNext(node, node.Left)
	}
	if node.Right != nil {
		connectCurrent(node.Right)
	}
	if node.Left != nil {
		connectCurrent(node.Left)
	}
}
