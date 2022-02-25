package task0202

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_isHappy(t *testing.T) {
	assert.True(t, isHappy(19))
}

func Test_isHappy2(t *testing.T) {
	assert.False(t, isHappy(2))
}

func Test_isHappy3(t *testing.T) {
	assert.True(t, isHappy(7))
}

func Test_isHappy4(t *testing.T) {
	assert.False(t, isHappy(3))
}
