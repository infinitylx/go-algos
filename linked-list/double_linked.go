package main

type DoubleLinkedList struct {
	Head *Node
	Tail *Node
	size int
}

type Node struct {
	Value    string
	Next     *Node
	Previous *Node
}

func (ll *DoubleLinkedList) AddFront(v string) {
	var node = new(Node)
	ll.Head.Previous = node
	node.Next = ll.Head
	node.Value = v
	ll.Head = node
	ll.size++
}

func (ll *DoubleLinkedList) AddBack(v string) {
	var node = new(Node)
	node.Value = v
	node.Previous = ll.Tail
	ll.Tail.Next = node
	ll.Tail = node
	ll.size++
}

func (ll *DoubleLinkedList) GetFirst() string {
	return ll.Head.Value
}

func (ll *DoubleLinkedList) GetLast() string {
	return ll.Tail.Value
}

func (ll *DoubleLinkedList) Size() int {
	return ll.size
}

func (ll *DoubleLinkedList) Clear() {
	ll.Head = nil
	ll.Tail = nil
}

// Delete - can be used in parallel from both direction. Size should be considered during using parallel computation.
func (ll *LinkedList) Ddelete(v string) *LinkedList {
	// rewrite this method to what???
	var nextNode = ll.Next
	if ll.Value == v {
		return ll.Next
	}
	for ll.Next != nil {
		if nextNode.Value == v {
			// remove this element
			ll.Next = nextNode.Next
			break
		}
		ll = ll.Next
		nextNode = ll.Next
	}
	return ll
}
