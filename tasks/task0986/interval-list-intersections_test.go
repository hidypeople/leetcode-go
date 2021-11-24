package task0986

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_intervalIntersection(t *testing.T) {
	first := [][]int{{0, 2}, {5, 10}, {13, 23}, {24, 25}}
	second := [][]int{{1, 5}, {8, 12}, {15, 24}, {25, 26}}
	assert.Equal(t,
		[][]int{{1, 2}, {5, 5}, {8, 10}, {15, 23}, {24, 24}, {25, 25}},
		intervalIntersection(first, second))
}

func Test_intervalIntersection2(t *testing.T) {
	first := [][]int{{0, 2}, {5, 10}, {13, 23}, {24, 25}}
	second := [][]int{}
	assert.Equal(t,
		[][]int{},
		intervalIntersection(first, second))
}

func Test_intervalIntersection3(t *testing.T) {
	first := [][]int{}
	second := [][]int{{1, 5}, {8, 12}, {15, 24}, {25, 26}}
	assert.Equal(t,
		[][]int{},
		intervalIntersection(first, second))
}

func Test_intervalIntersection4(t *testing.T) {
	first := [][]int{{1, 7}}
	second := [][]int{{3, 10}}
	assert.Equal(t,
		[][]int{{3, 7}},
		intervalIntersection(first, second))
}
