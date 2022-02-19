package task1675

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_minimumDeviation(t *testing.T) {
	assert.Equal(t, 1, minimumDeviation([]int{1, 2, 3, 4}))
}

func Test_minimumDeviation2(t *testing.T) {
	assert.Equal(t, 3, minimumDeviation([]int{4, 1, 5, 20, 3}))
}

func Test_minimumDeviation3(t *testing.T) {
	assert.Equal(t, 3, minimumDeviation([]int{2, 10, 8}))
}

func Test_minimumDeviation4(t *testing.T) {
	assert.Equal(t, 0, minimumDeviation([]int{8, 1, 2, 1}))
}
