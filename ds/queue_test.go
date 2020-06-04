package ds

import (
	"testing"
)

func TestPushOp(t *testing.T) {
	var queue Queue
	queue.Push(5)
	queue.Push(4)

	ele, _ := queue.pop()

	if ele != 5 && queue.size == 1 {
		t.Errorf("did not pop the correct element or the size is incorrect")
	}
}

func TestPop(t *testing.T) {
	var queue Queue
	ele, _ := queue.pop()
	if ele != nil || queue.size != 0 {
		t.Errorf("this list was empty so pop should have returned nil")
	}
}

func TestPeek(t *testing.T) {
	var queue Queue
	queue.Push(5)
	queue.Push(4)
	queue.Push(5)
	queue.Push(4)
	ele, _ := queue.Peek()
	if ele != 5 || queue.size != 4 {
		t.Errorf("peek should not have removed element from the queue")
	}
}

func TestIsEmpty(t *testing.T) {
	var queue Queue
	queue.Push(5)
	queue.Push(4)
	queue.Push(5)
	queue.Push(4)
	if queue.IsEmpty() {
		t.Errorf("queue is not empty, IsEmpty should have returned false")
	}
}
