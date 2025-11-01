package main

import (
	"errors"
	"fmt"
)

type Node struct {
	Data int
	Next *Node
}

type LinkedList struct {
	Head   *Node
	Tail   *Node
	Length int
}

func (ll *LinkedList) InsertAt(index, data int) error {
	if index < 0 || index >= ll.Length {
		return errors.New("index out of range")
	}

	newNode := Node{data, nil}

	if index == 0 {
		newNode.Next = ll.Head
		ll.Head = &newNode
	} else if index == ll.Length-1 {
		ll.Tail.Next = &newNode
		ll.Tail = &newNode
	} else {
		currentNode := ll.Head
		for range index - 1 {
			currentNode = currentNode.Next
		}

		newNode.Next = currentNode.Next
		currentNode.Next = &newNode
	}

	ll.Length++

	return nil
}

func (ll *LinkedList) AddToFront(data int) {
	newNode := Node{data, nil}
	if ll.Head == nil {
		ll.Head = &newNode
		ll.Tail = &newNode
	} else {
		newNode.Next = ll.Head
		ll.Head = &newNode
	}

	ll.Length++
}

func (ll *LinkedList) AddToBack(data int) {
	newNode := Node{data, nil}
	if ll.Head == nil {
		ll.Head = &newNode
		ll.Tail = &newNode
	} else {
		ll.Tail.Next = &newNode
		ll.Tail = &newNode
	}

	ll.Length++
}

func (ll *LinkedList) Remove(index int) error {
	if index < 0 || index >= ll.Length {
		return errors.New("index out of range")
	}

	if index == 0 {
		ll.Head = ll.Head.Next
		if ll.Head == nil {
			ll.Tail = nil
		}
	} else if index == ll.Length-1 {
		currentNode := ll.Head
		for range index - 1 {
			currentNode = currentNode.Next
		}
		currentNode.Next = nil
		ll.Tail = currentNode
	} else {
		currentNode := ll.Head
		for range index - 1 {
			currentNode = currentNode.Next
		}
		currentNode.Next = currentNode.Next.Next
	}

	ll.Length--
	return nil
}

func main() {
	linkedList := LinkedList{}
	for i := range 5 {
		if i == 0 {
			continue
		}
		linkedList.AddToBack(i)
	}

	currentNode := linkedList.Head
	for i := 0; i < linkedList.Length; i++ {
		fmt.Println(currentNode.Data)
		currentNode = currentNode.Next
	}

	fmt.Println()

	err := linkedList.Remove(1)
	if err != nil {
		panic(err)
	}

	currentNode = linkedList.Head
	for i := 0; i < linkedList.Length; i++ {
		fmt.Println(currentNode.Data)
		currentNode = currentNode.Next
	}

	err = linkedList.InsertAt(1, 90)
	if err != nil {
		panic(err)
	}

	fmt.Println()

	currentNode = linkedList.Head
	for i := 0; i < linkedList.Length; i++ {
		fmt.Println(currentNode.Data)
		currentNode = currentNode.Next
	}
}
