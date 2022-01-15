package task01413

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_minStartValue(t *testing.T) {
	assert.Equal(t, 5, minStartValue([]int{-3, 2, -3, 4, 2}))
}

func Test_minStartValue2(t *testing.T) {
	assert.Equal(t, 1, minStartValue([]int{1, 2}))
}

func Test_minStartValue3(t *testing.T) {
	assert.Equal(t, 5, minStartValue([]int{1, -2, -3}))
}
