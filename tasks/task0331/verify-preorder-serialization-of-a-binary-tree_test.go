package task0331

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_isValidSerialization0(t *testing.T) {
	assert.True(t, isValidSerialization("#"))
}

func Test_isValidSerialization(t *testing.T) {
	assert.True(t, isValidSerialization("9,3,4,#,#,1,#,#,2,#,6,#,#"))
}

func Test_isValidSerialization2(t *testing.T) {
	assert.False(t, isValidSerialization("1,#"))
}

func Test_isValidSerialization3(t *testing.T) {
	assert.False(t, isValidSerialization("9,#,#,1"))
}

func Test_isValidSerialization4(t *testing.T) {
	assert.True(t, isValidSerialization("9,#,92,#,#"))
}
