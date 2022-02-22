package task0122

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_maxProfit(t *testing.T) {
	assert.Equal(t, 7, maxProfit([]int{7, 1, 5, 3, 6, 4}))
}

func Test_maxProfit2(t *testing.T) {
	assert.Equal(t, 4, maxProfit([]int{1, 2, 3, 4, 5}))
}

func Test_maxProfit3(t *testing.T) {
	assert.Equal(t, 0, maxProfit([]int{7, 6, 4, 3, 1}))
}

func Test_maxProfit4(t *testing.T) {
	assert.Equal(t, 7, maxProfit([]int{3, 2, 6, 5, 0, 3}))
}

func Test_maxProfit5(t *testing.T) {
	assert.Equal(t, 18, maxProfit([]int{5, 5, 4, 9, 3, 8, 5, 5, 1, 6, 8, 3, 4}))
}

func Test_maxProfit6(t *testing.T) {
	assert.Equal(t, 3, maxProfit([]int{1, 2, 3, 4}))
}
