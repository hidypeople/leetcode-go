package task0402

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_removeKdigits(t *testing.T) {
	assert.Equal(t, "1219", removeKdigits("1432219", 3))
}

func Test_removeKdigits2(t *testing.T) {
	assert.Equal(t, "200", removeKdigits("10200", 1))
}

func Test_removeKdigits3(t *testing.T) {
	assert.Equal(t, "0", removeKdigits("10", 2))
}

func Test_removeKdigits4(t *testing.T) {
	assert.Equal(t, "11", removeKdigits("112", 1))
}

func Test_removeKdigits5(t *testing.T) {
	assert.Equal(t, "123", removeKdigits("12345", 2))
}
