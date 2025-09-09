package main

import (
	"errors"
	"fmt"
	"hash/fnv"
)

var ErrDuplicateKey = errors.New("duplicate key")

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
	fmt.Println(err)
}
