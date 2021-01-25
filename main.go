package main

import( 
	"github.com/sudhanshusngh/Stack/pkg/stack"
)

func main() {
	s := stack.Stack{}
	s.Push(2)
	s.Print()
	s.Push(5)
	s.Print()
	s.Push(7)
	s.Print()
	s.Push(9)
	s.Print()
	s.Push(4)
	s.Print()
	s.Pop()
	s.Print()
}
