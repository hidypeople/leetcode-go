package task0165

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_compareVersion(t *testing.T) {
	assert.Equal(t, 0, compareVersion("1.01", "1.00001"))
}

func Test_compareVersion2(t *testing.T) {
	assert.Equal(t, 0, compareVersion("1.0", "1.0.0"))
}

func Test_compareVersion3(t *testing.T) {
	assert.Equal(t, -1, compareVersion("0.1", "1.1"))
}

func Test_compareVersion4(t *testing.T) {
	assert.Equal(t, 0, compareVersion("1.0.000.0000", "1.0.0"))
}
