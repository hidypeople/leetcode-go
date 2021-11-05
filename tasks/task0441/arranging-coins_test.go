package task0441

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_arrangeCoins(t *testing.T) {
	assert.Equal(t, 2, arrangeCoins(5))
}

func Test_arrangeCoins2(t *testing.T) {
	assert.Equal(t, 3, arrangeCoins(8))
}

func Test_arrangeCoins3(t *testing.T) {
	assert.Equal(t, 1, arrangeCoins(1))
}

func Test_arrangeCoins5(t *testing.T) {
	assert.Equal(t, 5, arrangeCoins(15))
}
