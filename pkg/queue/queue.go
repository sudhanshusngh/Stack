package queue

import "fmt"

type Node struct {
	data int
	next *Node
}

type Queue struct {
	first *Node
	last *Node
}

func(q *Queue) PushQ(val int) {
	n := Node{}
	n.data = val
	if q.first == nil {
		q.first = &n
		q.last = q.first
		return
	}
	q.last.next = &n
	q.last = q.last.next
	return
}

func(q *Queue) PrintQueue() {
	ptr := q.first
	for ptr != nil {
		fmt.Printf("%d-> ", ptr.data)
		ptr = ptr.next
	}
	fmt.Println("NIL")
}


