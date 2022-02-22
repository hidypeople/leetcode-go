package task0168

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_convertToTitle(t *testing.T) {
	assert.Equal(t, "A", convertToTitle(1))
}

func Test_convertToTitle2(t *testing.T) {
	assert.Equal(t, "AB", convertToTitle(28))
}

func Test_convertToTitle3(t *testing.T) {
	assert.Equal(t, "ZY", convertToTitle(701))
}
