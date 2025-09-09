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

func (l *Linklist) Delete(d string) error {
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

func (l *Linklist) Print() {
	for n := l.head; n != nil; n = n.Next {
		fmt.Println(n.Data)
	}
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

	var err error
	err = l.Delete("data 1")

	l.Print()

	err = l.Delete("data 9")
	err = l.Delete("data 100")

	err = l.Delete("data 300")
	fmt.Println("err:", err)
	l.Print()
}
