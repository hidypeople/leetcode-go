package task0006

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_convert(t *testing.T) {
	assert.Equal(t, "PAHNAPLSIIGYIR", convert("PAYPALISHIRING", 3))
	assert.Equal(t, "PINALSIGYAHRPI", convert("PAYPALISHIRING", 4))
	assert.Equal(t, "A", convert("A", 1))
}
