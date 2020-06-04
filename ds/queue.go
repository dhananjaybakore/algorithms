package ds

import "fmt"

//FIFO
type Queue struct {
	top  *Node
	last *Node
	size int
}

func (queue *Queue) pop() (value interface{}, err error) {
	if queue.top != nil {
		value, queue.top = queue.top.data, queue.top.next
		queue.size--
		return value, nil
	}
	return nil, fmt.Errorf("queue is empty")
}

func (queue *Queue) Push(item interface{}) {
	last := &Node{item, nil}
	if queue.last != nil {
		queue.last.next = last
	}
	if queue.top == nil {
		queue.top = last
	}
	queue.last = last
	queue.size++
}

func (queue *Queue) Peek() (interface{}, error) {
	if !queue.IsEmpty() {
		return queue.top.data, nil
	}
	return nil, fmt.Errorf("queue is empty")
}

func (queue *Queue) IsEmpty() bool {
	return queue.top == nil
}
