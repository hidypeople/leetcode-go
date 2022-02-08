package task0258

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_addDigits(t *testing.T) {
	assert.Equal(t, 2, addDigits(38))
}

func Test_addDigits2(t *testing.T) {
	assert.Equal(t, 0, addDigits(0))
}

func Test_addDigits3(t *testing.T) {
	assert.Equal(t, 1, addDigits(1234))
}
