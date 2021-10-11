package task0043

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_multiply(t *testing.T) {
	assert.Equal(t, "0", multiply("10123", "0"))
}

func Test_multiply2(t *testing.T) {
	assert.Equal(t, "6", multiply("2", "3"))
}

func Test_multiply3(t *testing.T) {
	assert.Equal(t, "56088", multiply("123", "456"))
}

func Test_multiply4(t *testing.T) {
	assert.Equal(t, "1219326311370217952237463801111263526900", multiply("12345678901234567890", "98765432109876543210"))
}
