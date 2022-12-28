package binary

import (
	"testing"
)

func TestInsertFind(t *testing.T) {
	var keys = []int{3, 2, 5}
	var values = []string{"Three", "Two", "Five"}
	var bt = new(BTree)

	for i := 0; i < len(keys); i++ {
		if bt.Insert(keys[i], values[i]) != nil {
			t.Error("Get incorrect value!")
		}
	}
}

func TestFind(t *testing.T) {
	var keys = []int{3, 2, 5}
	var values = []string{"Three", "Two", "Five"}
	var bt = new(BTree)

	for i := 0; i < len(keys); i++ {
		if bt.Insert(keys[i], values[i]) != nil {
			t.Error("Get incorrect value!")
		}
	}
	for i := 0; i < len(keys); i++ {
		var v, _ = bt.Find(keys[i])
		if values[i] != v {
			t.Error("Get incorrect value!")
		}
	}

}

func TestDelete(t *testing.T) {
	var keys = []int{3, 2, 9, 5, 7, 1, 6}
	var values = []string{"Three", "Two", "Nine", "Five", "Seven", "One", "Six"}
	var bt = new(BTree)

	for i := 0; i < len(keys); i++ {
		if bt.Insert(keys[i], values[i]) != nil {
			t.Error("Get incorrect value!")
		}
	}

	if bt.Delete(3) != nil {
		t.Error("Failed delete!")
	}
	var _, err = bt.Find(3)
	if err == nil {
		t.Error(err)
	}
	_, err = bt.Find(5)
	if err != nil {
		t.Error(err)
	}
	_, err = bt.Find(6)
	if err != nil {
		t.Error(err)
	}

}
