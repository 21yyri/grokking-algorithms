package main

import "fmt"

type Node struct {
	Data int
	Next *Node
}

type Stack struct {
	Top *Node
}

func (s *Stack) Push(data int) {
	newNode := Node{data, nil}

	if s.Top == nil {
		s.Top = &newNode
	} else {
		newNode.Next = s.Top
		s.Top = &newNode
	}
}

func (s *Stack) Pop() {
	if s.Top == nil {
		return
	}

	s.Top = s.Top.Next
}

func (s *Stack) Peek() int {
	if s.Top == nil {
		return -1
	}

	return s.Top.Data
}

func main() {
	stack := Stack{}

	for i := range 5 {
		stack.Push(i)
	}

	stack.Pop()

	fmt.Println(stack.Peek())
}
