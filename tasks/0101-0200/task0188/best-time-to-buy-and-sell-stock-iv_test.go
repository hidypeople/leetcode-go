package task0188

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_maxProfit(t *testing.T) {
	assert.Equal(t, 2, maxProfit(2, []int{2, 4, 1}))
}

func Test_maxProfit2(t *testing.T) {
	assert.Equal(t, 7, maxProfit(2, []int{3, 2, 6, 5, 0, 3}))
}
