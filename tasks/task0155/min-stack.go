package task0155

// Design a stack that supports push, pop, top, and retrieving the minimum element in constant time.
// Implement the MinStack class:
//   MinStack() initializes the stack object.
//   void push(int val) pushes the element val onto the stack.
//   void pop() removes the element on the top of the stack.
//   int top() gets the top element of the stack.
//   int getMin() retrieves the minimum element in the stack.
//   type MinStack struct {}
//
// Constraints:
//   -2^31 <= val <= 2^31 - 1
//   Methods pop, top and getMin operations will always be called on non-empty stacks.
//   At most 3 * 10^4 calls will be made to push, pop, top, and getMin.
//
// Results:
//   Runtime: 16 ms, faster than 97.88% of Go online submissions for Min Stack.
//   Memory Usage: 8.5 MB, less than 76.72% of Go online submissions for Min Stack.
type MinStack struct {
	n     int
	mins  []int
	items []int
}

func Constructor() MinStack {
	return MinStack{
		n:     0,
		mins:  make([]int, 0),
		items: make([]int, 0),
	}
}

func (s *MinStack) Push(val int) {
	min := val
	if s.n > 0 && s.mins[s.n-1] < val {
		min = s.mins[s.n-1]
	}
	s.items = append(s.items, val)
	s.mins = append(s.mins, min)
	s.n++
}

func (s *MinStack) Pop() {
	s.items = s.items[:s.n-1]
	s.mins = s.mins[:s.n-1]
	s.n--
}

func (s *MinStack) Top() int {
	return s.items[s.n-1]
}

func (s *MinStack) GetMin() int {
	return s.mins[s.n-1]
}
