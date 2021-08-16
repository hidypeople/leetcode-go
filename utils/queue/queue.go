package queue

// Queue interface
type IQueue interface {
	Len() int
	Enqueue(x int)
	Dequeue() (int, bool)
}

// Queue struct
type Queue struct {
	queue []int
}

// Create new empty queue
func QueueFactory() *Queue {
	return &Queue{
		queue: make([]int, 0),
	}
}

// Current queue size
func (t *Queue) Len() int {
	return len(t.queue)
}

// Add item into the queue
func (t *Queue) Enqueue(x int) {
	t.queue = append(t.queue, x)
}

// Get item from the queue
func (t *Queue) Dequeue() (int, bool) {
	if len(t.queue) == 0 {
		return 0, false
	}
	x := t.queue[0]
	t.queue = t.queue[1:]
	return x, true
}
