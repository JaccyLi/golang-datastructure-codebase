package doublylinkedlist

type DoublyLinkedList struct {
	Size int64
	Head *Node
	Tail *Node
}

type Node struct {
	DataStore interface{}
	Prev      *Node
	Next      *Node
}

// behavior implementations
//
