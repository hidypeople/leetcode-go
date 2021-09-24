package combination

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsEqualCombinations(t *testing.T) {
	assert.Equal(t, true, IsEqualCombinations(
		[][]int{},
		[][]int{},
		true))
}

func TestIsEqualCombinations2(t *testing.T) {
	assert.Equal(t, false, IsEqualCombinations(
		[][]int{{1}},
		[][]int{{2}},
		true))
}

func TestIsEqualCombinations3(t *testing.T) {
	assert.Equal(t, true, IsEqualCombinations(
		[][]int{{1}, {2}},
		[][]int{{2}, {1}},
		true))
	assert.Equal(t, true, IsEqualCombinations(
		[][]int{{1}, {2}},
		[][]int{{1}, {2}},
		true))
}

func TestIsEqualCombinations4(t *testing.T) {
	assert.Equal(t, true, IsEqualCombinations(
		[][]int{{1}, {1, 1, 1, 1}, {2, 2}},
		[][]int{{1}, {2, 2}, {1, 1, 1, 1}},
		true))
	assert.Equal(t, false, IsEqualCombinations(
		[][]int{{1}, {1, 2, 1, 1}, {2, 2}},
		[][]int{{1}, {2, 2}, {1, 1, 1, 1}},
		true))
}

func TestIsEqualCombinations5(t *testing.T) {
	assert.Equal(t, true, IsEqualCombinations(
		[][]int{{4, 4, 4, 4}, {}, {1}, {1, 4, 4}, {1, 4}, {1, 4, 4, 4}, {1, 4, 4, 4, 4}, {4}, {4, 4}, {4, 4, 4}},
		[][]int{{}, {1}, {1, 4}, {1, 4, 4}, {1, 4, 4, 4}, {1, 4, 4, 4, 4}, {4}, {4, 4}, {4, 4, 4}, {4, 4, 4, 4}},
		true))
}

func TestIsEqualCombinations_OrderSensitiveSetting(t *testing.T) {
	assert.Equal(t, false, IsEqualCombinations(
		[][]int{{1}, {1, 1, 1, 1}, {2, 1}},
		[][]int{{1}, {1, 2}, {1, 1, 1, 1}},
		true))

	assert.Equal(t, true, IsEqualCombinations(
		[][]int{{1}, {1, 1, 1, 1}, {2, 1}},
		[][]int{{1}, {1, 2}, {1, 1, 1, 1}},
		false))
}
