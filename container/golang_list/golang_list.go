package main

import (
	"container/list"
	"fmt"
)

func main() {
	doubleLinkedList := list.New()

	doubleLinkedList.PushBack("banana")
	doubleLinkedList.PushBack("orange")
	doubleLinkedList.PushFront("apple")

	fmt.Printf("double linked list front: %v\n", doubleLinkedList.Front().Value)
	fmt.Printf("double linked list back: %v\n", doubleLinkedList.Back().Value)
}
