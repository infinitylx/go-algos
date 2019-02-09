package main

import (
	"testing"
)

func TestPutGet(t *testing.T) {
	var ht = New()
	var keys = []string{"key1", "key2", "key3"}
	var values = []string{"value1", "value2", "value3"}
	for i := 0; i < len(keys); i++ {
		ht.Put(keys[i], values[i])
	}

	for i := 0; i < len(keys); i++ {
		tvalue, err := ht.Get(keys[i])
		if err != nil {
			t.Error("Get don't work.")
			t.Error(err)
		}
		if tvalue != values[i] {
			t.Error("Get incorrect value!")
			t.Error(tvalue)
		}
	}
}

func TestRemove(t *testing.T) {
	var ht = New()
	var keys = []string{"key1", "key2", "key3"}
	var values = []string{"value1", "value2", "value3"}
	for i := 0; i < len(keys); i++ {
		ht.Put(keys[i], values[i])
	}

	ht.Remove(keys[1])
	_, err := ht.Get(keys[1])

	if err == nil {
		t.Error("Get non nil error.")
	}
	err = ht.Remove(keys[1])

	if err == nil {
		t.Error("Get non nil error.")
	}
}

func TestCollision(t *testing.T) {
	// put multiple keys under one index
	// check get
	// check remove
}
