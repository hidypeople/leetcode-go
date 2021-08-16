package queue

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestQueue_Enqueue(t *testing.T) {
	queue := QueueFactory()
	assert.Equal(t, 0, queue.Len())
}

func TestQueue_Enqueue2(t *testing.T) {
	queue := QueueFactory()
	assert.Equal(t, 0, queue.Len())
	queue.Enqueue(1)
	val, ok := queue.Dequeue()
	assert.True(t, ok)
	assert.Equal(t, 1, val)
	assert.Equal(t, 0, queue.Len())
	_, ok = queue.Dequeue()
	assert.False(t, ok)
}
