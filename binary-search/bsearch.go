package binary

import "errors"

// TODO: rewrite the whole structure.

type BTree struct {
	head *BTreeNode
	id   int // <- wtf?
}

type BTreeNode struct {
	left   *BTreeNode
	right  *BTreeNode
	parent *BTreeNode
	value  string
	key    int
}

func (bt *BTree) Insert(key int, value string) error {
	if bt.head == nil {
		bt.head = new(BTreeNode)
		bt.head.key = key
		bt.head.value = value
		return nil
	}
	bt.head.insert(key, value)
	return nil
}

func (bt *BTree) Find(key int) (string, error) {
	if bt.head == nil {
		return "", errors.New("no such key")
	}
	return bt.head.find(key)
}

func (bt *BTree) Delete(key int) error {
	if bt.head == nil {
		return errors.New("no such key")
	}
	return bt.head.delete(key)
}

func (btn *BTreeNode) insert(key int, value string) error {
	if key == btn.key {
		btn.value = value
	} else {
		if key < btn.key {
			if btn.left == nil {
				newNode := new(BTreeNode)
				newNode.key = key
				newNode.value = value
				newNode.parent = btn
				btn.left = newNode
			} else {
				return btn.left.insert(key, value)
			}
		} else if key > btn.key {
			if btn.right == nil {
				newNode := new(BTreeNode)
				newNode.key = key
				newNode.value = value
				newNode.parent = btn
				btn.right = newNode
			} else {
				return btn.right.insert(key, value)
			}
		}
	}
	return nil
}

func (btn *BTreeNode) findNode(key int) (r *BTreeNode, err error) {
	// DON'T USE THIS STYLE WITH NAMED RETURN IN F-SIGNATURE
	// var r *BTreeNode
	// var err error
	if key == btn.key {
		r = btn
	} else {
		if key < btn.key {
			if btn.left == nil {
				return nil, errors.New("no such key")
			}
			r, err = btn.left.findNode(key)
		}
		if key > btn.key {
			if btn.right == nil {
				return nil, errors.New("no such key")
			}
			r, err = btn.right.findNode(key)
		}
	}
	return // r, err
}

func (btn *BTreeNode) find(key int) (string, error) {
	var r string
	var err error
	if key == btn.key {
		return btn.value, nil
	} else {
		if key < btn.key {
			if btn.left == nil {
				return "", errors.New("no such key")
			}
			r, err = btn.left.find(key)
		}
		if key > btn.key {
			if btn.right == nil {
				return "", errors.New("no such key")
			}
			r, err = btn.right.find(key)
		}
	}
	return r, err
}

func (btn *BTreeNode) isLeaf() bool {
	if btn.left == nil && btn.right == nil {
		return true
	} else {
		return false
	}
}

func (btn *BTreeNode) min() *BTreeNode {
	if btn.left == nil {
		return btn
	} else {
		return btn.left.min()
	}
}

// code of total bullshit. delete has btn pointer and operates on node??? wtf???
func (btn *BTreeNode) delete(key int) error {
	var node, err = btn.findNode(key)
	if err != nil {
		return err
	}
	// node is leaf. simply remove link from parent node
	if node.isLeaf() {
		if node.parent.left == node {
			node.parent.left = nil
		} else {
			node.parent.right = nil
		}
		return nil
	}
	// node has both child
	if node.left != nil && node.right != nil {
		// find min in right sub tree.
		// replace node with it.
		// if min node have child procede with delete procedure for node with one child.
		var t = node.right.min()
		node.key = t.key
		node.value = t.value
		node.right.delete(t.key)
		return nil

	} else {
		// has only one child
		var child *BTreeNode
		if node.left != nil {
			child = node.left
		} else {
			child = node.right
		}
		if node.parent.left == node {
			node.parent.left = child
		} else {
			node.parent.right = child
		}
		child.parent = node.parent
		return nil
	}

	return errors.New("error")
}
