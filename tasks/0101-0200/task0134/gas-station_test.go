package task0134

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_canCompleteCircuit(t *testing.T) {
	gas := []int{1, 2, 3, 4, 5}
	cost := []int{3, 4, 5, 1, 2}
	assert.Equal(t, 3, canCompleteCircuit(gas, cost))
}

func Test_canCompleteCircuit2(t *testing.T) {
	gas := []int{2, 3, 4}
	cost := []int{3, 4, 3}
	assert.Equal(t, -1, canCompleteCircuit(gas, cost))
}

func Test_canCompleteCircuit3(t *testing.T) {
	gas := []int{5, 1, 2, 3, 4}
	cost := []int{4, 4, 1, 5, 1}
	assert.Equal(t, 4, canCompleteCircuit(gas, cost))
}

func Test_canCompleteCircuit4(t *testing.T) {
	gas := []int{5, 8, 2, 8}
	cost := []int{6, 5, 6, 6}
	assert.Equal(t, 3, canCompleteCircuit(gas, cost))
}
