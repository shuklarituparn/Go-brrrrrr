package main

/*
type Stack struct {
	elements []int
}

func (s *Stack) Push(element int) {
	s.elements = append(s.elements, element)
}

func (s *Stack) Pop() int {
	if len(s.elements) == 0 {
		panic("Stack is empty")
	}
	lastIndex := len(s.elements) - 1
	popped := s.elements[lastIndex]
	s.elements = s.elements[:lastIndex]
	return popped
}

func (s *Stack) Peek() int {
	if len(s.elements) == 0 {
		panic("Stack is empty")
	}
	lastIndex := len(s.elements) - 1
	return s.elements[lastIndex]
}

func (s *Stack) IsEmpty() bool {
	return len(s.elements) == 0
}
*/
//second implementation of stack in go
