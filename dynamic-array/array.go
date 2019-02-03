package main

import "errors"

type DArray struct {
	Size     int
	Capacity int
	Content  []string
}

func MakeDArray(size int, capacity int) DArray {
	var content = make([]string, size, capacity)
	return DArray{size, capacity, content}
}

func (d *DArray) newCap() int {
	var newCap int
	if d.Capacity <= d.Size {
		newCap = d.Capacity * 2
	} else {
		newCap = d.Capacity
	}
	d.Capacity = newCap
	return newCap
}

func (d *DArray) Insert(i int, v string) error {
	if !(i <= d.Size) {
		return errors.New("index out of range")
	}
	var newContent = make([]string, d.Size+1, d.newCap())

	if i > 0 {
		copy(newContent, d.Content[0:i])
	}
	newContent[i] = v
	for index := i; index < d.Size; index++ {
		newContent[index+1] = d.Content[index]
	}
	d.Content = newContent
	d.Size++
	return nil
}

func (d *DArray) Get(i int) (string, error) {
	if !(i <= d.Size) {
		return "", errors.New("index out of range")
	}
	return d.Content[i], nil
}

func (d *DArray) Set(i int, v string) error {
	if !(i < d.Size) {
		return errors.New("index out of range")
	}
	d.Content[i] = v
	return nil
}

func (d *DArray) Delete(i int) error {
	if !(i < d.Size) {
		return errors.New("index out of range")
	}
	var newContent []string
	ns := d.Size - 1
	newContent = make([]string, ns, d.Capacity)
	if i > 0 {
		copy(newContent, d.Content[0:i])
	}
	for index := i; index < d.Size-1; index++ {
		newContent[index] = d.Content[index+1]
	}
	d.Content = newContent
	d.Size = ns
	return nil
}

func (d *DArray) IsEmpty() bool {
	return d.Size == 0
}

func (d *DArray) Contains(v string) bool {
	for _, el := range d.Content {
		if v == el {
			return true
		}
	}
	return false
}
