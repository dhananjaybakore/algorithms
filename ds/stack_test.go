package ds

import (
	"testing"
)

func TestPushAndPeekOperation(t *testing.T) {
	var stack Stack
	stack.Push(2)
	stack.Push(3)
	stack.Push(4)
	stack.Push(5)
	ele := stack.Peek()
	if ele != 5 && !stack.IsEmpty() {
		t.Errorf("elements not inserted properly to the stack, top ele should be:5 but got %v", ele)
	}
	if stack.size != 4 {
		t.Errorf("stack size should have been 4")
	}

}

func TestPopOperation(t *testing.T) {
	var stack Stack
	stack.Push(2)
	stack.Push(3)
	stack.Push(4)
	stack.Push(5)
	topEle := stack.Peek()
	if topEle != 5 && !stack.IsEmpty() {
		t.Errorf("the popped element should be: %v but got:%v", 5, topEle)
	}
	ele := stack.Pop()
	if ele != 5 && !stack.IsEmpty() {
		t.Errorf("the top element should be: %v but got: %v", 5, ele)
	}
}

func TestPopOnOneElementStack(t *testing.T) {
	var stack Stack
	stack.Push(2)
	stack.Pop()
}

func TestPopEmptyStack(t *testing.T) {
	var stack Stack
	stack.Pop()
}
