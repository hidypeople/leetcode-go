package task0115

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_numDistinct(t *testing.T) {
	assert.Equal(t, 3, numDistinct("rabbbit", "rabbit"))
}

func Test_numDistinct2(t *testing.T) {
	assert.Equal(t, 5, numDistinct("babgbag", "bag"))
}
