package main

import (
	"errors"
	"fmt"
)

var (
	ErrNotFound = errors.New("item not found")
)

type Linklist struct {
	head *node
	tail *node
	size int
}

func (l *Linklist) Add(d string) {
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

func (l *Linklist) Del(d string) error {
	if l.head == nil {
		return fmt.Errorf("%w %v", ErrNotFound, d)
	}
	if l.head.Data == d {
		l.head = l.head.Next
		l.size--
		if l.head == nil {
			l.tail = nil
		}
		return nil
	}
	prev := l.head
	curr := l.head.Next
	for curr != nil {
		if curr.Data == d {
			prev.Next = curr.Next
			if curr == l.tail {
				l.tail = prev
			}
			l.size--
			return nil
		}
		prev = curr
		curr = curr.Next
	}
	return fmt.Errorf("%w %v", ErrNotFound, d)
}

func (l *Linklist) Len() int {
	return l.size
}

func (l *Linklist) String() string {
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

func (l *Linklist) Get(idx int) (string, error) {
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

func main() {
	fmt.Println("linklist")
	l := new(Linklist)
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

	err = l.Del("data 1")

	fmt.Println(l)

	err = l.Del("data 9")
	err = l.Del("data 100")

	err = l.Del("data 300")
	fmt.Println("err:", err)
	fmt.Println(l)

	fmt.Println(l.Len())
}
