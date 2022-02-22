package task0123

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_maxProfit(t *testing.T) {
	assert.Equal(t, 6, maxProfit([]int{3, 3, 5, 0, 0, 3, 1, 4}))
}

func Test_maxProfit5(t *testing.T) {
	assert.Equal(t, 6, maxProfit([]int{3, 5, 0, 3, 1, 4}))
}

func Test_maxProfit2(t *testing.T) {
	assert.Equal(t, 4, maxProfit([]int{1, 2, 3, 4, 5}))
}

func Test_maxProfit3(t *testing.T) {
	assert.Equal(t, 0, maxProfit([]int{7, 6, 4, 3, 1}))
}

func Test_maxProfit4(t *testing.T) {
	assert.Equal(t, 0, maxProfit([]int{1}))
}

func Test_maxProfit6(t *testing.T) {
	assert.Equal(t, 7, maxProfit([]int{6, 1, 3, 2, 4, 7}))
}

func Test_maxProfit7(t *testing.T) {
	assert.Equal(t, 13, maxProfit([]int{1, 2, 4, 2, 5, 7, 2, 4, 9, 0}))
}
