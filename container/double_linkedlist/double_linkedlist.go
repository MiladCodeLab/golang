package main

import (
	"errors"
	"fmt"
	"iter"
)

var (
	ErrNotFound = errors.New("item not found")
)

type DoubleLinkedlist struct {
	head *node
	tail *node
	size int
}

func (l *DoubleLinkedlist) Append(d string) {
	n := new(node)
	n.Data = d
	l.size++
	if l.head == nil {
		l.head = n
		l.tail = n
		return
	}
	l.tail.next = n
	n.prev = l.tail
	l.tail = n
}

func (l *DoubleLinkedlist) Prepend(d string) {
	n := new(node)
	n.Data = d
	l.size++
	if l.head == nil {
		l.head, l.tail = n, n
		return
	}
	n.next = l.head
	l.head.prev = n
	l.head = n
}

func (l *DoubleLinkedlist) Pop() (string, error) {
	if l.tail == nil {
		return "", ErrNotFound
	}
	var val node
	val = *l.tail

	if l.tail == l.head {
		l.tail = nil
		l.head = nil
	} else {
		l.tail = l.tail.prev
		l.tail.next = nil
	}
	l.size--
	return val.Data, nil
}

func (l *DoubleLinkedlist) Shift() (string, error) {
	if l.head == nil {
		return "", ErrNotFound
	}
	var val node
	val = *l.head
	if l.head == l.tail {
		l.head = nil
		l.tail = nil
	} else {
		l.head = l.head.next
		l.head.prev = nil
	}
	l.size--
	return val.Data, nil
}

func (l *DoubleLinkedlist) Del(idx int) error {
	if l.head == nil {
		return ErrNotFound
	}

	if idx < 0 || idx >= l.size {
		return fmt.Errorf("%w: %d", ErrNotFound, idx)
	}

	var curr *node
	if idx < l.size/2 {
		curr = l.head
		for i := 0; i < idx; i++ {
			curr = curr.next
		}
	} else {
		curr = l.tail
		for i := l.size - 1; i > idx; i-- {
			curr = curr.prev
		}
	}

	if curr.prev != nil {
		curr.prev.next = curr.next
	} else {
		l.head = curr.next
	}

	if curr.next != nil {
		curr.next.prev = curr.prev
	} else {
		l.tail = curr.prev
	}
	l.size--

	return nil
}

func (l *DoubleLinkedlist) Len() int {
	return l.size
}

func (l *DoubleLinkedlist) String() string {
	var (
		items string
		idx   int
	)

	for n := l.head; n != nil; n = n.next {
		items += fmt.Sprintf("#%d: %v\t", idx, n.Data)
		idx++
		if idx%3 == 0 {
			items += "\n"
		}
	}
	return items
}

func (l *DoubleLinkedlist) Get(idx int) (string, error) {
	if idx < 0 || idx >= l.size {
		return "", fmt.Errorf("%w %v", ErrNotFound, idx)
	}

	curr := l.head
	for i := 0; i < idx; i++ {
		curr = curr.next
	}
	return curr.Data, nil
}

type node struct {
	Data string
	next *node
	prev *node
}

func (l *DoubleLinkedlist) Iter() iter.Seq[string] {
	return func(yield func(string) bool) {
		for n := l.head; n != nil; n = n.next {
			if !yield(n.Data) {
				return
			}
		}
	}
}

func (l *DoubleLinkedlist) Iter2() iter.Seq2[int, string] {
	return func(yield func(int, string) bool) {
		idx := 0
		for n := l.head; n != nil; n = n.next {
			if !yield(idx, n.Data) {
				return
			}
			idx++
		}
	}
}

func main() {
	l := new(DoubleLinkedlist)
	l.Append("data 1")
	l.Append("data 5")
	l.Append("data 2")
	l.Append("data 9")
	l.Append("data 100")

	var (
		val string
		err error
	)
	fmt.Println("*****Get*****")
	val, err = l.Get(0)
	fmt.Println(val, err)

	val, err = l.Get(3)
	fmt.Println(val, err)

	val, err = l.Get(4)
	fmt.Println(val, err)

	fmt.Println("*****Prepend*****")
	l.Prepend("banana")
	l.Prepend("apple")

	fmt.Println("*****Del*****")

	err = l.Del(1)
	fmt.Println(err)
	err = l.Del(5)
	fmt.Println(err)

	fmt.Println("*****Len*****")
	fmt.Println(l.Len())

	fmt.Println("*****Iter*****")
	for data := range l.Iter() {
		fmt.Println(data)

	}

	fmt.Println("*****Iter2*****")
	for i, data := range l.Iter2() {
		fmt.Println(i, " ", data)
	}
}
