package task0380

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRandomizedSet_Insert(t *testing.T) {
	randomizedSet := Constructor()
	assert.True(t, randomizedSet.Insert(1))
	assert.True(t, randomizedSet.Insert(2))
	assert.False(t, randomizedSet.Insert(2))
}

func TestRandomizedSet_Remove(t *testing.T) {
	randomizedSet := Constructor()
	assert.True(t, randomizedSet.Insert(1))
	assert.True(t, randomizedSet.Insert(2))
	assert.False(t, randomizedSet.Remove(10000))
	assert.True(t, randomizedSet.Remove(1))
	assert.True(t, randomizedSet.Remove(2))
	assert.False(t, randomizedSet.Remove(2))
}

func TestRandomizedSet_GetRandom(t *testing.T) {
	randomizedSet := Constructor()
	randomizedSet.Insert(1)
	randomizedSet.Insert(2)
	assert.Contains(t, []int{1, 2}, randomizedSet.GetRandom())
	assert.Contains(t, []int{1, 2}, randomizedSet.GetRandom())
	assert.Contains(t, []int{1, 2}, randomizedSet.GetRandom())
	randomizedSet.Insert(3)
	assert.Contains(t, []int{1, 2, 3}, randomizedSet.GetRandom())
	assert.Contains(t, []int{1, 2, 3}, randomizedSet.GetRandom())
	assert.Contains(t, []int{1, 2, 3}, randomizedSet.GetRandom())
	assert.Contains(t, []int{1, 2, 3}, randomizedSet.GetRandom())
}
