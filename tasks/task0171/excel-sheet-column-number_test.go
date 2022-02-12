package task0171

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_titleToNumber(t *testing.T) {
	assert.Equal(t, 1, titleToNumber("A"))
}

func Test_titleToNumber2(t *testing.T) {
	assert.Equal(t, 28, titleToNumber("AB"))
}

func Test_titleToNumber3(t *testing.T) {
	assert.Equal(t, 701, titleToNumber("ZY"))
}
