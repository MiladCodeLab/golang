package main

import (
	"errors"
	"fmt"
	"hash/fnv"
	"iter"
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
	list     []*list
	size     int
	capacity int
}

func NewHashMap(capacity int) *HashMap {
	return &HashMap{
		list:     make([]*list, capacity),
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

	if h.list[idx] == nil {
		h.list[idx] = &list{
			head: newNode,
		}
		h.size++
		return nil
	}

	curr := h.list[idx].head
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
	if h.list[idx] == nil {
		return fmt.Errorf("%w: %s", ErrNotFound, key)
	}
	curr := h.list[idx].head
	var prev *node
	for curr != nil {
		if curr.key == key {
			if prev == nil {
				h.list[idx].head = curr.next
			} else {
				prev.next = curr.next
			}
			h.size--
			return nil
		}

		prev = curr
		curr = curr.next
	}
	return fmt.Errorf("%w: %s", ErrNotFound, key)
}

func (h *HashMap) Get(key string) (string, error) {
	idx := int(h.hashFunction(key)) % h.capacity
	if h.list[idx] == nil {
		return "", fmt.Errorf("%w: %s", ErrNotFound, key)
	}

	curr := h.list[idx].head
	for curr != nil {
		if curr.key == key {
			return curr.value, nil
		}
		curr = curr.next
	}
	return "", fmt.Errorf("%w: %s", ErrNotFound, key)
}
func (h *HashMap) Len() int {
	return h.size
}

func (h *HashMap) Iter() iter.Seq2[string, string] {
	return func(yield func(string, string) bool) {
		for _, v := range h.list {
			if v == nil {
				continue
			}
			for n := v.head; n != nil; n = n.next {
				if !yield(n.key, n.value) {
					return
				}
			}
		}
	}
}

func (h *HashMap) String() string {
	var vals string
	for i, v := range h.list {
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

	fmt.Println("*****Append*****")
	err = mp.Add("fruit22", "apple")
	err = mp.Add("fruit169", "banana")
	err = mp.Add("fruit94", "cucumber")
	err = mp.Add("fruit281", "orange")
	fmt.Println(mp)

	fmt.Println("*****Iter*****")
	for key, val := range mp.Iter() {
		fmt.Printf("key: %s, value: %s\n", key, val)
	}

	fmt.Println("*****Get*****")
	var val string
	val, err = mp.Get("fruit94")
	fmt.Println(val, " err:", err)

	fmt.Println("*****Del*****")
	err = mp.Del("fruit169")
	fmt.Println(mp, " ", err)
	err = mp.Del("fruit22")
	fmt.Println(mp, " ", err)
	err = mp.Del("fruit281")
	fmt.Println(mp, " ", err)

	fmt.Println("*****Len*****")
	fmt.Println(mp.Len())

}
