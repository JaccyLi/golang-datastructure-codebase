package tree

// Node contains a tree's node
type Node struct {
	Key int
	Left *Node
	Right *Node
}

// Insert insert a node to a tree
func (n *Node)Insert(k int) {
	if n.Key > k {
		// put left
		if n.Left == nil {
			// put
			n.Left = &Node{Key: k}
		} else {
			// recursively call self
			n.Left.Insert(k)
		}
	} else if n.Key < k {
		// put right
		if n.Right == nil {
			// put
			n.Right = &Node{Key: k}
		} else {
			// recursively call self
			n.Right.Insert(k)
		}
	}
}

// Search search a node in a tree
func (n *Node) Search(k int) bool {
	if n == nil {
		return false
	}

	if n.Key > k {
		// call again
		return n.Left.Search(k)
	} else if n.Key < k {
		// call again
		return n.Right.Search(k)
	}

	return true
}