package main

import( 
	"github.com/sudhanshusngh/Stack/pkg/stack"
	"github.com/sudhanshusngh/Stack/pkg/queue"
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
	q := queue.Queue{}
	q.PushQ(3)
	q.PushQ(9)
	q.PushQ(5)
	q.PushQ(7)
	q.PushQ(2)
	q.PushQ(4)
	q.PrintQueue()
}
