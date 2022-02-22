package task0119

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_getRow(t *testing.T) {
	assert.Equal(t, []int{1, 3, 3, 1}, getRow(3))
}

func Test_getRow2(t *testing.T) {
	assert.Equal(t, []int{1}, getRow(0))
}

func Test_getRow3(t *testing.T) {
	assert.Equal(t, []int{1, 1}, getRow(1))
}

func Test_getRow4(t *testing.T) {
	assert.Equal(t, []int{1, 5, 10, 10, 5, 1}, getRow(5))
}

func Test_getRow5(t *testing.T) {
	assert.Equal(t, []int{1, 6, 15, 20, 15, 6, 1}, getRow(6))
}

func Test_getRow6(t *testing.T) {
	assert.Equal(t, 34, len(getRow(33)))
}

func Test_getRow7(t *testing.T) {
	assert.Equal(t,
		[]int{1, 30, 435, 4060, 27405, 142506, 593775, 2035800, 5852925, 14307150, 30045015, 54627300, 86493225, 119759850, 145422675, 155117520, 145422675, 119759850, 86493225, 54627300, 30045015, 14307150, 5852925, 2035800, 593775, 142506, 27405, 4060, 435, 30, 1},
		getRow(30))
}

func Test_getRow8(t *testing.T) {
	assert.Equal(t, []int{1, 6, 15, 20, 15, 6, 1}, getRow2(6))
}
