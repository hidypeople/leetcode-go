package task0169

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_majorityElement(t *testing.T) {
	assert.Equal(t, 3, majorityElement([]int{3, 2, 3}))
}

func Test_majorityElement2(t *testing.T) {
	assert.Equal(t, 2, majorityElement([]int{2, 2, 1, 1, 1, 2, 2}))
}
