package stack

// Stackable interface
type IStack interface {
	Len() int
	Push(x int)
	Pop() (int, bool)
}

// Simple stack class
type Stack struct {
	stack []int
}

// Stack factory: creates new empty stacks
func StackFactory() *Stack {
	return &Stack{
		stack: make([]int, 0),
	}
}

// Current stack size
func (t *Stack) Len() int {
	return len(t.stack)
}

// Push item into the stack
func (t *Stack) Push(x int) {
	t.stack = append(t.stack, x)
}

// Pop item from the stack
func (t *Stack) Pop() (int, bool) {
	if t.Len() == 0 {
		return 0, false
	}
	x := t.stack[t.Len()-1]
	t.stack = t.stack[:t.Len()-1]
	return x, true
}
