package task0980

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_uniquePathsIII(t *testing.T) {
	assert.Equal(t, 4, uniquePathsIII([][]int{
		{1, 0, 0, 0},
		{0, 0, 0, 0},
		{0, 0, 0, 2}}))
}

func Test_uniquePathsIII2(t *testing.T) {
	assert.Equal(t, 0, uniquePathsIII([][]int{
		{0, 1},
		{2, 0}}))
}

func Test_uniquePathsIII3(t *testing.T) {
	assert.Equal(t, 1, uniquePathsIII([][]int{
		{1, 0, 0},
		{2, 0, 0}}))
}
