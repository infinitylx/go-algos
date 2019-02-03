package main

type LinkedList struct {
	Value string
	Next  *LinkedList
}

func (ll *LinkedList) AddFront(v string) *LinkedList {
	var node = new(LinkedList)
	node.Next = ll
	node.Value = v
	// ll = node
	return node
}

func (ll *LinkedList) AddBack(v string) {
	var node = new(LinkedList)
	node.Value = v
	var tll = ll
	for ; tll.Next != nil; tll = tll.Next {
	}
	tll.Next = node
}

func (ll *LinkedList) GetFirst() string {
	return ll.Value
}

func (ll *LinkedList) GetLast() string {
	if ll.Next == nil {
		return ll.Value
	} else {
		return ll.Next.GetLast()
	}
}

func (ll *LinkedList) Size() int {
	var r = 1
	for ; ll.Next != nil; ll = ll.Next {
		r++
	}
	return r
}

func (ll *LinkedList) Clear() *LinkedList {
	return new(LinkedList)
}

func (ll *LinkedList) Delete(v string) *LinkedList {
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
