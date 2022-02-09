package task0166

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_fractionToDecimal(t *testing.T) {
	assert.Equal(t, "0.5", fractionToDecimal(1, 2))
}

func Test_fractionToDecimal2(t *testing.T) {
	assert.Equal(t, "2", fractionToDecimal(2, 1))
}

func Test_fractionToDecimal3(t *testing.T) {
	assert.Equal(t, "0.(012)", fractionToDecimal(4, 333))
}

func Test_fractionToDecimal4(t *testing.T) {
	assert.Equal(t, "0.1(6)", fractionToDecimal(1, 6))
}

func Test_fractionToDecimal5(t *testing.T) {
	assert.Equal(t, "3.(142857)", fractionToDecimal(22, 7))
}

func Test_fractionToDecimal6(t *testing.T) {
	assert.Equal(t, "0.0000000004656612873077392578125", fractionToDecimal(-1, -2147483648))
}

func Test_fractionToDecimal7(t *testing.T) {
	assert.Equal(t, "-2", fractionToDecimal(2, -1))
	assert.Equal(t, "-2", fractionToDecimal(-2, 1))
}
