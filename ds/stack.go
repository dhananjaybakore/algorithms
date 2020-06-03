package ds

type Stack struct {
	top  *Node
	size int
}

type Node struct {
	data interface{}
	next *Node
}

func (stack *Stack) Pop() (value interface{}) {
	if stack.size > 0 {
		value, stack.top = stack.top.data, stack.top.next
		stack.size--
		return value
	}
	return nil
}

func (stack *Stack) Push(item interface{}) {
	stack.top = &Node{item, stack.top}
	stack.size++
}

func (stack *Stack) Peek() interface{} {
	return stack.top.data
}

func (stack *Stack) IsEmpty() bool {
	return stack.top == nil
}
