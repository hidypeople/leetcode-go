package task0474

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_findMaxForm(t *testing.T) {
	strs := []string{"10", "0001", "111001", "1", "0"}
	m := 5
	n := 3
	assert.Equal(t, 4, findMaxForm(strs, m, n))
}

func Test_findMaxForm2(t *testing.T) {
	strs := []string{"10", "0", "1"}
	m := 1
	n := 1
	assert.Equal(t, 2, findMaxForm(strs, m, n))
}
