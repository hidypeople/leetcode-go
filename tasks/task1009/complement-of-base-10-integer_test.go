package task1009

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_bitwiseComplement(t *testing.T) {
	assert.Equal(t, 2, bitwiseComplement(5))
}

func Test_bitwiseComplement2(t *testing.T) {
	assert.Equal(t, 0, bitwiseComplement(7))
}

func Test_bitwiseComplement3(t *testing.T) {
	assert.Equal(t, 5, bitwiseComplement(10))
}
