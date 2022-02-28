package task0228

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_summaryRanges(t *testing.T) {
	assert.Equal(t, []string{"0->2", "4->5", "7"}, summaryRanges([]int{0, 1, 2, 4, 5, 7}))
}

func Test_summaryRanges2(t *testing.T) {
	assert.Equal(t, []string{"0", "2->4", "6", "8->9"}, summaryRanges([]int{0, 2, 3, 4, 6, 8, 9}))
}
