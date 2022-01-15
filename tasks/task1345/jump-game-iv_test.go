package task1345

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_minJumps(t *testing.T) {
	assert.Equal(t, 3, minJumps([]int{100, -23, -23, 404, 100, 23, 23, 23, 3, 404}))
}

func Test_minJumps2(t *testing.T) {
	assert.Equal(t, 0, minJumps([]int{7}))
}

func Test_minJumps3(t *testing.T) {
	assert.Equal(t, 1, minJumps([]int{7, 6, 9, 6, 9, 6, 9, 7}))
}

func Test_minJumps4(t *testing.T) {
	// 11 -> 11 -> 22 -> 22 -> 33 -> 33
	assert.Equal(t, 5, minJumps([]int{
		11,
		-101, -102, -103, -104, -105, -106, -107, -108, -109, -110,
		33, 22,
		-201, -202, -203, -204, -205, -206, -207, -208, -209, -210,
		22, 11,
		-401, -402, -403, -404, -405, -406, -407, -408, -409, -410,
		33,
	}))
}
