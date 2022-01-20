package task0875

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_minEatingSpeed(t *testing.T) {
	assert.Equal(t, 4, minEatingSpeed([]int{3, 6, 7, 11}, 8))
}

func Test_minEatingSpeed2(t *testing.T) {
	assert.Equal(t, 30, minEatingSpeed([]int{30, 11, 23, 4, 20}, 5))
}

func Test_minEatingSpeed3(t *testing.T) {
	assert.Equal(t, 23, minEatingSpeed([]int{30, 11, 23, 4, 20}, 6))
}

func Test_minEatingSpeed4(t *testing.T) {
	assert.Equal(t, 2, minEatingSpeed([]int{312884470}, 312884469))
}
