package task0065

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_isNumber(t *testing.T) {
	validNumbers := []string{".1", "2", "0089", "-0.1", "+3.14", "4.", "-.9", "2e10", "-90E3", "3e+7", "+6e-1", "53.5e93", "-123.456e789"}
	for _, number := range validNumbers {
		assert.True(t, isNumber(number))
	}
}

func Test_isNumber2(t *testing.T) {
	invalidNumbers := []string{"+E3", ".", ".e1", "4e+", "123.123.33", "1e+123+12", "abc", "1a", "1e", "e3", "99e2.5", "--6", "-+3", "95a54e53"}
	for _, number := range invalidNumbers {
		assert.False(t, isNumber(number))
	}
}
