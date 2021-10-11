package task0076

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_minWindow(t *testing.T) {
	assert.Equal(t, "BANC", minWindow("ADOBECODEBANC", "ABC"))
}

func Test_minWindow2(t *testing.T) {
	assert.Equal(t, "a", minWindow("a", "a"))
}

func Test_minWindow3(t *testing.T) {
	assert.Equal(t, "", minWindow("a", "aa"))
}

func Test_minWindow4(t *testing.T) {
	assert.Equal(t, "BxAxC", minWindow("xAxBxBxAxCx", "ABC"))
}

func Test_minWindow5(t *testing.T) {
	assert.Equal(t, "ABxC", minWindow("xAxBxCxxxxABxCx", "ABC"))
}
