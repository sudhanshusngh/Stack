package stack

import "fmt"

type node struct{
	data int
	next *node
}

type Stack struct {
	top *node
}

func (s *Stack) Push(val int) {
	n := node{}
	n.data = val
	if s.top == nil {
		s.top = &n
		return
	}
	n.next = s.top
	s.top = &n
	return
}

func (s *Stack) Pop() int {
	if s.top == nil {
		return -1
	}
	ptr := s.top
	s.top = s.top.next
	ptr.next = nil
	return ptr.data
}

func (s *Stack) Print() {
	ptr := s.top
	for ptr != nil {
		fmt.Printf("%d->", ptr.data)
		ptr = ptr.next
	}
	fmt.Println("NIL")
}
