package task0093

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_restoreIpAddresses(t *testing.T) {
	assert.Equal(t, []string{"255.255.11.135", "255.255.111.35"}, restoreIpAddresses("25525511135"))
}

func Test_restoreIpAddresses2(t *testing.T) {
	assert.Equal(t, []string{"0.0.0.0"}, restoreIpAddresses("0000"))
}

func Test_restoreIpAddresses3(t *testing.T) {
	assert.Equal(t, []string{"1.1.1.1"}, restoreIpAddresses("1111"))
}

func Test_restoreIpAddresses4(t *testing.T) {
	assert.Equal(t, []string{"0.10.0.10", "0.100.1.0"}, restoreIpAddresses("010010"))
}

func Test_restoreIpAddresses5(t *testing.T) {
	assert.Equal(t, []string{"1.0.10.23", "1.0.102.3", "10.1.0.23", "10.10.2.3", "101.0.2.3"}, restoreIpAddresses("101023"))
}

func Test_restoreIpAddresses6(t *testing.T) {
	assert.Equal(t, []string{}, restoreIpAddresses("00000"))
}
