package task0038

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_countAndSay(t *testing.T) {
	assert.Equal(t, "1", countAndSay(1))
}

func Test_countAndSay2(t *testing.T) {
	assert.Equal(t, "1211", countAndSay(4))
}
