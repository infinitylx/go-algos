package main

import (
	"errors"
	"hash/fnv"
)

const DEFAULTCAPACITY = 16
const NOSUCHKEYERROR = "No such key"

type HashTable struct {
	capacity uint32
	storage  []*KeyValueList
}

type KeyValueList struct {
	keyvalue [2]string
	next     *KeyValueList
}

func New() *HashTable {
	var ht = new(HashTable)
	ht.capacity = DEFAULTCAPACITY
	ht.storage = make([]*KeyValueList, DEFAULTCAPACITY)
	return ht
}

func (ht *HashTable) Put(key, value string) {
	// check size and capacity if needed resize
	// calculate hash and index
	var index = ht.getIndex(getHash(key))
	if ht.storage[index] == nil {
		ht.storage[index] = new(KeyValueList)
	}
	ht.storage[index].put(key, value)
}

func (ht *HashTable) Get(key string) (string, error) {
	var index = ht.getIndex(getHash(key))
	if ht.storage[index] == nil {
		return "", errors.New(NOSUCHKEYERROR)
	}
	return ht.storage[index].get(key)
}

func (ht *HashTable) Remove(key string) error {
	var index = ht.getIndex(getHash(key))
	if ht.storage[index] == nil {
		return errors.New(NOSUCHKEYERROR)
	}
	if ht.storage[index].keyvalue[0] == key {
		ht.storage[index] = ht.storage[index].next
		return nil
	}
	ht.storage[index].remove(key)
	return nil
}

func (kvl *KeyValueList) put(key, value string) {
	if kvl.keyvalue[0] == "" {
		// empty first element
		kvl.keyvalue[0] = key
		kvl.keyvalue[1] = value
		return
	}
	if kvl.keyvalue[0] == key {
		kvl.keyvalue[1] = value
		return
	}
	// check collision
	var nv *KeyValueList
	for ; nv != nil; nv = nv.next {
		if nv.keyvalue[0] == key {
			nv.keyvalue[1] = value
			return
		}
	}
	// add new collision
	var nkvl = new(KeyValueList)
	nkvl.put(key, value)
	nv.next = nkvl
}

func (kvl *KeyValueList) get(key string) (string, error) {
	if kvl.keyvalue[0] == key {
		return kvl.keyvalue[1], nil
	}
	for t := kvl; t.next != nil; t = t.next {
		if t.keyvalue[0] == key {
			return t.keyvalue[1], nil
		}
	}
	return "", errors.New(NOSUCHKEYERROR)
}

func (kvl *KeyValueList) remove(key string) error {
	var found = false
	for t := kvl; t.next != nil; t = t.next {
		if t.next.keyvalue[0] == key {
			if t.next.next != nil {
				t.next = t.next.next
			} else {
				t.next = nil
			}
			found = true
			break
		}
	}
	if found {
		return nil
	}
	return errors.New(NOSUCHKEYERROR)
}

func getHash(k string) uint32 {
	h := fnv.New32a()
	h.Write([]byte(k))
	return h.Sum32()
}

func (ht *HashTable) getIndex(kh uint32) uint32 {
	return kh % ht.capacity
}

func (ht *HashTable) resize(nc uint32) {
	// resize internal storage to new capacity

}
