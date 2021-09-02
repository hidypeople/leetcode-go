package task0050

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_myPow(t *testing.T) {
	assert.Equal(t, 1.0, myPow(123, 0))
}

func Test_myPow2(t *testing.T) {
	assert.Equal(t, 123.0, myPow(123, 1))
}

func Test_myPow3(t *testing.T) {
	assert.Equal(t, -123.0, myPow(-123, 1))
}

func Test_myPow4(t *testing.T) {
	assert.Equal(t, 1024.0, myPow(2.00000, 10))
}

func Test_myPow5(t *testing.T) {
	assert.Equal(t, 0.25, myPow(2.00000, -2))
}

func Test_myPow6(t *testing.T) {
	assert.True(t, abs(9.261-myPow(2.1, 3)) < 0.00001)
}

func abs(a float64) float64 {
	if a < 0 {
		return -a
	}
	return a
}
