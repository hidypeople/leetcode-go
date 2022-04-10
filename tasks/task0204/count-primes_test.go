package task0204

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_countPrimes(t *testing.T) {
	assert.Equal(t, 4, countPrimes(10))
}

func Test_countPrimes2(t *testing.T) {
	assert.Equal(t, 0, countPrimes(0))
}

func Test_countPrimes3(t *testing.T) {
	assert.Equal(t, 0, countPrimes(1))
}

func Test_countPrimes4(t *testing.T) {
	assert.Equal(t, 41538, countPrimes(499999))
}
