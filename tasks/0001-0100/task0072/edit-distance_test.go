package task0072

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_minDistance(t *testing.T) {
	assert.Equal(t, 0, minDistance("", ""))
}

func Test_minDistance2(t *testing.T) {
	assert.Equal(t, 4, minDistance("aaaa", ""))
}

func Test_minDistance3(t *testing.T) {
	assert.Equal(t, 4, minDistance("", "aaaa"))
}

func Test_minDistance4(t *testing.T) {
	assert.Equal(t, 3, minDistance("horse", "ros"))
}

func Test_minDistance5(t *testing.T) {
	assert.Equal(t, 5, minDistance("intention", "execution"))
}
