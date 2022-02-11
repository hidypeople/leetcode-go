package task0567

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_checkInclusion(t *testing.T) {
	assert.True(t, checkInclusion("ab", "eidbaooo"))
}

func Test_checkInclusion2(t *testing.T) {
	assert.False(t, checkInclusion("ab", "eidboaoo"))
}
