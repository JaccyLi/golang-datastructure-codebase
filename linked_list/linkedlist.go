package linkedlist

type LinkedList struct {
	Head *Node
}

type Node struct {
	DataStore interface{}
	Next      *Node
}

// behavior implementations
//
