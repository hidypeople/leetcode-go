package task0067

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_addBinary(t *testing.T) {
	assert.Equal(t, "100", addBinary("11", "1"))
}

func Test_addBinary2(t *testing.T) {
	assert.Equal(t, "10101", addBinary("1010", "1011"))
}

func Test_addBinary3(t *testing.T) {
	assert.Equal(t, "10000000", addBinary("1111111", "1"))
}

func Test_addBinary4(t *testing.T) {
	assert.Equal(t, "1110", addBinary("111", "111"))
}

func Test_addBinary5(t *testing.T) {
	assert.Equal(t, "0001", addBinary("0001", "0"))
}
