package task0134

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_singleNumber(t *testing.T) {
	assert.Equal(t, 1, singleNumber([]int{2, 2, 1}))
}

func Test_singleNumber2(t *testing.T) {
	assert.Equal(t, 4, singleNumber([]int{4, 1, 2, 1, 2}))
}

func Test_singleNumber3(t *testing.T) {
	assert.Equal(t, 1, singleNumber([]int{1}))
}
