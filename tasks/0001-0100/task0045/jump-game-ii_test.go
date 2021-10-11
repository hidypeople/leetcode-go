package task0045

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_jump(t *testing.T) {
	assert.Equal(t, 2, jump([]int{2, 3, 1, 1, 4}))
}

func Test_jump2(t *testing.T) {
	assert.Equal(t, 2, jump([]int{2, 3, 0, 1, 4}))
}

func Test_jump3(t *testing.T) {
	assert.Equal(t, 0, jump([]int{1}))
}

func Test_jump4(t *testing.T) {
	assert.Equal(t, 5, jump([]int{5, 6, 4, 4, 6, 9, 4, 4, 7, 4, 4, 8, 2, 6, 8, 1, 5, 9, 6, 5, 2, 7, 9, 7, 9, 6, 9, 4, 1, 6, 8, 8, 4, 4, 2, 0, 3, 8, 5}))
}

func Test_jump5(t *testing.T) {
	assert.Equal(t, 1, jump([]int{1, 2}))
}

func Test_jump6(t *testing.T) {
	assert.Equal(t, 3, jump([]int{1, 1, 2, 1, 1}))
}
