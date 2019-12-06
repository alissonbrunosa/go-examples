package stack

import "fmt"

type Stack struct {
	elements []interface{}
}

func (s *Stack) Push(element interface{}) {
	s.elements = append(s.elements, element)
}

func (s *Stack) Pop() interface{} {
	top := s.elements[len(s.elements)-1]
	s.elements = s.elements[0 : s.Length()-1]
	return top
}

func (s *Stack) Peek() interface{} {
	return s.elements[len(s.elements)-1]
}

func (s *Stack) Print() {
	for i := len(s.elements); i >= 0; i-- {
		fmt.Printf("Value is %v \n", s.elements[i])
	}
}

func (s *Stack) IsEmpty() bool {
	return s.Length() == 0
}

func (s *Stack) Length() int {
	return len(s.elements)
}
