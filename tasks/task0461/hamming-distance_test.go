package task0461

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_hammingDistance(t *testing.T) {
	assert.Equal(t, 2, hammingDistance(1, 4))
}

func Test_hammingDistance2(t *testing.T) {
	assert.Equal(t, 1, hammingDistance(1, 3))
}

func Test_hammingDistance3(t *testing.T) {
	assert.Equal(t, 1, hammingDistance(3, 1))
}
