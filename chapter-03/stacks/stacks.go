package main

import "fmt"

type Node struct {
	Data string
	Next *Node
}

type Stack struct {
	Top *Node
}

func (s *Stack) Push(data string) {
	newNode := Node{data, nil}

	if s.Top == nil {
		s.Top = &newNode
	} else {
		newNode.Next = s.Top
		s.Top = &newNode
	}
}

func (s *Stack) Pop() string {
	if s.Top == nil {
		return ""
	}

	poppedItem := s.Top
	s.Top = s.Top.Next

	return poppedItem.Data
}

func (s *Stack) Peek() string {
	if s.Top == nil {
		return ""
	}

	return s.Top.Data
}

func (s *Stack) isEmpty() bool {
	return s.Top == nil
}

func main() {
	stack := Stack{}
	strToReverse := "Hello, world!"

	result := ""

	for i := 0; i < len(strToReverse); i++ {
		stack.Push(string(strToReverse[i]))
	}

	for stack.Top != nil {
		result += stack.Pop()
	}

	fmt.Println(result)
}
