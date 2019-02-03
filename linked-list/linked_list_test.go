package main

import "testing"

func TestAddFront(t *testing.T) {
	var ll = new(LinkedList)
	ll.Value = "first"
	ll = ll.AddFront("head")
	if ll.GetFirst() != "head" {
		t.Error("AddFront don't work.")
	}
}

func TestAddBack(t *testing.T) {
	var ll = new(LinkedList)
	ll.Value = "first"
	ll.AddBack("last2")
	ll.AddBack("last1")
	if ll.GetLast() != "last1" {
		t.Error("AddFront don't work.")
	}
}

func TestGetFirst(t *testing.T) {
	var ll = LinkedList{"first", nil}
	if ll.GetFirst() != "first" {
		t.Error("GetFirst don't work.")
	}
}

func TestGetLast(t *testing.T) {
	var ll = new(LinkedList)
	ll.Value = "first"
	ll = ll.AddFront("head")
	if ll.GetLast() != "first" {
		t.Error("GetFirst don't work.")
	}
}

func TestSize(t *testing.T) {
	var ll = new(LinkedList)
	ll.Value = "first"
	if ll.Size() != 1 {
		t.Error("Size don't match. Should be 1 got ", ll.Size())
	}
	ll = ll.AddFront("head")
	ll.AddBack("last")

	if ll.Size() != 3 {
		t.Error("Size don't match. Should be 3 got ", ll.Size())
	}
}

func TestDelete(t *testing.T) {
	var ll = new(LinkedList)
	ll.Value = "first"
	ll = ll.AddFront("head")
	ll.AddBack("back")
	ll = ll.Delete("head")

	if ll.GetFirst() != "first" {
		t.Error("Delete don't work.")
	}
}
