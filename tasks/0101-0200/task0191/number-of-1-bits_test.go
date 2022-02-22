package task0191

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_hammingWeight(t *testing.T) {
	assert.Equal(t, 3, hammingWeight(0b00000000_00000000_00000000_00001011))
}

func Test_hammingWeight2(t *testing.T) {
	assert.Equal(t, 1, hammingWeight(0b00000000_00000000_00000000_10000000))
}

func Test_hammingWeight3(t *testing.T) {
	assert.Equal(t, 31, hammingWeight(0b11111111_11111111_11111111_11111101))
}
