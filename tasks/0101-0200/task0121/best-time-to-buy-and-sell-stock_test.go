package task0121

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_maxProfit(t *testing.T) {
	assert.Equal(t, 5, maxProfit([]int{7, 1, 5, 3, 6, 4}))
}

func Test_maxProfit2(t *testing.T) {
	assert.Equal(t, 0, maxProfit([]int{7, 6, 4, 3, 1}))
}
