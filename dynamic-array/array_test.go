package main

import "testing"

func TestGetInsert(t *testing.T) {
	t.Log("Creating dynamic-array...")
	var da = MakeDArray(1, 2)
	if da.Size != 1 {
		t.Error("Wrong size.")
	}
	ts := "first test string"
	ts2 := "second test string"
	ts3 := "third test string"
	ts4 := "third test string"
	tss := []string{ts, ts2, ts3, ts4}

	for i := 0; i < 3; i++ {
		err := da.Insert(0, tss[2-i])
		if err != nil {
			t.Error("Can't insert!")
		}
	}
	ierr := da.Insert(3, ts4)
	if ierr != nil {
		t.Error("Can't insert ts4")
	}

	for index := 0; index < 4; index++ {
		r, err := da.Get(index)
		if err != nil {
			t.Error("Can't get")
		}
		if r != tss[index] {
			t.Error("Got wrong value!", r)
		}
	}
}

func TestGetSet(t *testing.T) {
	t.Log("Creating dynamic-array...")
	var da = MakeDArray(5, 10)
	ts := "first test string"
	da.Set(3, ts)
	r, gerr := da.Get(3)
	if gerr != nil {
		t.Error("Can't get.")
	}
	if r != ts {
		t.Error("Get wrong value")
	}
}

func TestDelete(t *testing.T) {
	t.Log("Creating dynamic-array...")
	var da = MakeDArray(5, 10)
	ts := "first test string"
	da.Set(0, ts)
	derr := da.Delete(0)
	if derr != nil {
		t.Error("Can't delete.")
	}
	if da.Size != 4 {
		t.Error("Delete don't work!")
	}
}

func TestIsEmpty(t *testing.T) {
	var da = MakeDArray(3, 10)
	ts := "first test string"
	ts2 := "second test string"
	ts3 := "third test string"
	if da.IsEmpty() {
		t.Error("IsEmpty does not work!1")
	}
	if da.Set(0, ts) != nil {
		t.Error("Can't set ts")
	}
	if da.Set(1, ts2) != nil {
		t.Error("Can't set ts2")
	}
	if da.Set(2, ts3) != nil {
		t.Error("Can't set ts3")
	}
	if da.IsEmpty() {
		t.Error("IsEmpty does not work!2")
	}
	da.Delete(2)
	da.Delete(1)
	da.Delete(0)
	if !da.IsEmpty() {
		t.Error("IsEmpty does not work!3")
	}
}

func TestContains(t *testing.T) {
	var da = MakeDArray(5, 10)
	ts := "first test string"
	if da.Contains(ts) {
		t.Error("Contains doesn't work!")
	}
	da.Set(0, ts)
	if !da.Contains(ts) {
		t.Error("Contains doesn't work!")
	}
	da.Delete(0)
	if da.Contains(ts) {
		t.Error("Contains doesn't work!")
	}
}
