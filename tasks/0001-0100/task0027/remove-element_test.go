package task0026

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_removeElement(t *testing.T) {
	array := []int{3, 2, 2, 3}
	result := removeElement(array, 3)
	assert.Equal(t, 2, result)
	assert.Equal(t, []int{2, 2}, array[:result])
}

func Test_removeElement2(t *testing.T) {
	array := []int{0, 1, 2, 2, 3, 0, 4, 2}
	result := removeElement(array, 2)
	assert.Equal(t, 5, result)
	assert.Equal(t, []int{0, 1, 3, 0, 4}, array[:result])
}

func Test_removeElement3(t *testing.T) {
	array := []int{}
	result := removeElement(array, 2)
	assert.Equal(t, 0, result)
	assert.Equal(t, []int{}, array)
}
