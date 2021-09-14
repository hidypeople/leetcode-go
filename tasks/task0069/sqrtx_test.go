package task0069

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_mySqrt(t *testing.T) {
	assert.Equal(t, 2, mySqrt(4))
}

func Test_mySqrt2(t *testing.T) {
	assert.Equal(t, 2, mySqrt(8))
}

func Test_mySqrt3(t *testing.T) {
	assert.Equal(t, 0, mySqrt(0))
}

func Test_mySqrt4(t *testing.T) {
	assert.Equal(t, 4, mySqrt(24))
}

func Test_mySqrt5(t *testing.T) {
	assert.Equal(t, 23, mySqrt(555))
}
