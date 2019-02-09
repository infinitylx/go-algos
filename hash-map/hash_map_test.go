package main

import "testing"

func TestPutGet(t *testing.T) {
	var ht = New()
	var key = "key1"
	var value = "value1"
	ht.Put(key, value)

	tvalue, err := ht.Get(key)
	if err != nil {
		t.Error("Get don't work.")
		t.Error(err)
	}
	if tvalue != value {
		t.Error("Get incorrect value!")
		t.Error(tvalue)
	}
}

func TestRemove(t *testing.T) {
	var ht = New()
	var key = "key1"
	var value = "value1"
	ht.Put(key, value)

	ht.Remove(key)
	_, err := ht.Get(key)

	if err == nil {
		t.Error("Get non nil error.")
	}
}
