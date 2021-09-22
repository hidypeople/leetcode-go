package task0085

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_maximalRectangle(t *testing.T) {
	assert.Equal(t, 6, maximalRectangle([][]byte{
		{'1', '0', '1', '0', '0'},
		{'1', '0', '1', '1', '1'},
		{'1', '1', '1', '1', '1'},
		{'1', '0', '0', '1', '0'}}))
}

func Test_maximalRectangle2(t *testing.T) {
	assert.Equal(t, 0, maximalRectangle([][]byte{{'0'}}))
}

func Test_maximalRectangle3(t *testing.T) {
	assert.Equal(t, 0, maximalRectangle([][]byte{{'0', '0', '0'}}))
}

func Test_maximalRectangle4(t *testing.T) {
	assert.Equal(t, 1, maximalRectangle([][]byte{{'0', '1', '0'}}))
}

func Test_maximalRectangle5(t *testing.T) {
	assert.Equal(t, 0, maximalRectangle([][]byte{}))
}

func Test_maximalRectangle6(t *testing.T) {
	assert.Equal(t, 10, maximalRectangle([][]byte{
		{'0', '1', '1', '0', '0', '1', '0', '1', '0', '1'},
		{'0', '0', '1', '0', '1', '0', '1', '0', '1', '0'},
		{'1', '0', '0', '0', '0', '1', '0', '1', '1', '0'},
		{'0', '1', '1', '1', '1', '1', '1', '0', '1', '0'},
		{'0', '0', '1', '1', '1', '1', '1', '1', '1', '0'},
		{'1', '1', '0', '1', '0', '1', '1', '1', '1', '0'},
		{'0', '0', '0', '1', '1', '0', '0', '0', '1', '0'},
		{'1', '1', '0', '1', '1', '0', '0', '1', '1', '1'},
		{'0', '1', '0', '1', '1', '0', '1', '0', '1', '1'}}))
}
