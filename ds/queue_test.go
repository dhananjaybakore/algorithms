package ds

import "fmt"

//FIFO
type Queue struct {
	top  *Node
	last *Node
	size int
}

func (queue *Queue) pop() (value interface{}) {
	if queue.size > 0 {
		value, queue.top = queue.top.data, queue.top.next
		queue.size--
		return value
	}
	return nil
}

func (queue *Queue) Push(item interface{}) {
	last := &Node{item, nil}
	queue.last.next = last
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
