package main

import (
	"errors"
	"fmt"
	"hash/fnv"
)

var ErrDuplicateKey = errors.New("duplicate key")
var ErrNotFound = errors.New("entry not found")

type node struct {
	key   string
	value string
	next  *node
}
type list struct {
	head *node
}
type HashMap struct {
	elements []*list
	size     int
	capacity int
}

func NewHashMap(capacity int) *HashMap {
	return &HashMap{
		elements: make([]*list, capacity),
		size:     0,
		capacity: capacity,
	}
}

func (h *HashMap) Add(key string, value string) error {
	idx := int(h.hashFunction(key)) % h.capacity
	newNode := &node{
		key:   key,
		value: value,
		next:  nil,
	}

	if h.elements[idx] == nil {
		h.elements[idx] = &list{
			head: newNode,
		}
		h.size++
		return nil
	}

	curr := h.elements[idx].head
	for curr != nil {
		if curr.key == key {
			return fmt.Errorf("%w: %s", ErrDuplicateKey, curr.key)
		}

		if curr.next == nil {
			curr.next = newNode
			h.size++
			return nil
		}

		curr = curr.next
	}
	return nil
}

func (h *HashMap) Del(key string) error {
	idx := int(h.hashFunction(key)) % h.capacity
	if h.elements[idx] == nil {
		return fmt.Errorf("%w: %s", ErrNotFound, key)
	}
	curr := h.elements[idx].head
	prev := h.elements[idx].head
	for curr != nil {
		if curr.key == key {
			//Delete here
			prev.next = curr.next
			return nil
		}

		prev = curr
		curr = curr.next
	}
	return fmt.Errorf("%w: %s", ErrNotFound, key)
}

func (h *HashMap) String() string {
	var vals string
	for i, v := range h.elements {
		if v == nil {
			continue
		}
		for n := v.head; n != nil; n = n.next {
			vals += fmt.Sprintf("#%d: key: %s, value: %s\n", i, n.key, n.value)
		}
	}
	return vals
}

func (_ *HashMap) hashFunction(key string) uint32 {
	h := fnv.New32a()
	h.Write([]byte(key))
	return h.Sum32()
}
func main() {
	var err error
	mp := NewHashMap(100)
	err = mp.Add("key8", "is a decent software developer")
	err = mp.Add("key22", "is a decent animator")
	err = mp.Add("key149", "is a decent man")
	fmt.Println(mp)

	err = mp.Del("key22")
	fmt.Println(mp)
	err = mp.Del("key8")
	fmt.Println(mp)
	err = mp.Del("key149")
	fmt.Println(mp)

	fmt.Println(err)
}
