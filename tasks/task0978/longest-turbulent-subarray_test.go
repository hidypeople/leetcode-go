package task0978

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_maxTurbulenceSize(t *testing.T) {
	assert.Equal(t, 5, maxTurbulenceSize([]int{9, 4, 2, 10, 7, 8, 8, 1, 9}))
}

func Test_maxTurbulenceSize2(t *testing.T) {
	assert.Equal(t, 2, maxTurbulenceSize([]int{4, 8, 12, 16}))
}

func Test_maxTurbulenceSize3(t *testing.T) {
	assert.Equal(t, 1, maxTurbulenceSize([]int{1}))
}

func Test_maxTurbulenceSize4(t *testing.T) {
	assert.Equal(t, 8, maxTurbulenceSize([]int{0, 8, 45, 88, 48, 68, 28, 55, 17, 24}))
}
