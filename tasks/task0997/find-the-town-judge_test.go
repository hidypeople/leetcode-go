package task0997

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_findJudge(t *testing.T) {
	assert.Equal(t, 2, findJudge(2, [][]int{{1, 2}}))
}

func Test_findJudge2(t *testing.T) {
	assert.Equal(t, 3, findJudge(3, [][]int{{1, 3}, {2, 3}}))
}

func Test_findJudge3(t *testing.T) {
	assert.Equal(t, -1, findJudge(3, [][]int{{1, 3}, {2, 3}, {3, 1}}))
}

func Test_findJudge4(t *testing.T) {
	assert.Equal(t, 1, findJudge(1, [][]int{}))
}

func Test_findJudge5(t *testing.T) {
	assert.Equal(t, -1, findJudge(3, [][]int{{1, 2}, {2, 3}}))
}
