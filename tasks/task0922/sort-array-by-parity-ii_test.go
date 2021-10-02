package task0922

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_sortArrayByParityII(t *testing.T) {
	assert.Equal(t, []int{4, 5, 2, 7}, sortArrayByParityII([]int{4, 2, 5, 7}))
}

func Test_sortArrayByParityII2(t *testing.T) {
	assert.Equal(t, []int{2, 3}, sortArrayByParityII([]int{2, 3}))
}

func Test_sortArrayByParityII3(t *testing.T) {
	assert.Equal(t, []int{2, 1, 2, 3, 4, 3, 4, 5}, sortArrayByParityII([]int{1, 2, 2, 3, 3, 4, 4, 5}))
}
