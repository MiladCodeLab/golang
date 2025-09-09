package main

import (
	"errors"
	"fmt"
	"iter"
)

var (
	ErrNotFound = errors.New("item not found")
)

type Linkedlist struct {
	head *node
	tail *node
	size int
}

func (l *Linkedlist) Add(d string) {
	n := new(node)
	n.Data = d
	l.size++
	if l.head == nil {
		l.head = n
		l.tail = n
		return
	}
	l.tail.Next = n
	l.tail = n
}

func (l *Linkedlist) Del(idx int) error {
	if l.head == nil {
		return ErrNotFound
	}
	if idx < 0 || idx >= l.size {
		return fmt.Errorf("%w %v", ErrNotFound, idx)
	}
	if idx == 0 { // case 1
		l.head = l.head.Next
		l.size--
		if l.head == nil {
			l.tail = nil
		}
		return nil
	}

	prev := l.head
	for i := 0; i < idx-1; i++ {
		prev = prev.Next
	}
	curr := prev.Next
	prev.Next = curr.Next

	if curr == l.tail {
		l.tail = prev
	}
	l.size--
	return nil
}

func (l *Linkedlist) Len() int {
	return l.size
}

func (l *Linkedlist) String() string {
	var (
		items string
		idx   int
	)

	for n := l.head; n != nil; n = n.Next {
		items += fmt.Sprintf("#%d: %v\t", idx, n.Data)
		idx++
		if idx%3 == 0 {
			items += "\n"
		}
	}
	return items
}

func (l *Linkedlist) Get(idx int) (string, error) {
	if idx < 0 || idx >= l.size {
		return "", fmt.Errorf("%w %v", ErrNotFound, idx)
	}

	curr := l.head
	for i := 0; i < idx; i++ {
		curr = curr.Next
	}
	return curr.Data, nil
}

type node struct {
	Data string
	Next *node
}

func (l *Linkedlist) Iter() iter.Seq[string] {
	return func(yield func(string) bool) {
		for n := l.head; n != nil; n = n.Next {
			if !yield(n.Data) {
				return
			}
		}
	}
}

func (l *Linkedlist) Iter2() iter.Seq2[int, string] {
	return func(yield func(int, string) bool) {
		idx := 0
		for n := l.head; n != nil; n = n.Next {
			if !yield(idx, n.Data) {
				return
			}
			idx++
		}
	}
}

func main() {
	fmt.Println("linkedlist")
	l := new(Linkedlist)
	l.Add("data 1")
	l.Add("data 5")
	l.Add("data 2")
	l.Add("data 9")
	l.Add("data 100")

	var (
		val string
		err error
	)

	val, err = l.Get(0)
	fmt.Println(val, err)

	val, err = l.Get(3)
	fmt.Println(val, err)

	val, err = l.Get(4)
	fmt.Println(val, err)

	err = l.Del(1)

	fmt.Println(l)

	err = l.Del(0)
	err = l.Del(2)

	fmt.Println("err:", err)
	fmt.Println(l)

	fmt.Println(l.Len())

	// iterator

	for data := range l.Iter() {
		fmt.Println("inside Iter: ", data)

	}

	for i, data := range l.Iter2() {
		fmt.Println("inside Iter2: ", i, " ", data)

	}
}
