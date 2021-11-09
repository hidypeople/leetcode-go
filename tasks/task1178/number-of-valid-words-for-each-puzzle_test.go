package task1178

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_findNumOfValidWords(t *testing.T) {
	assert.Equal(t, []int{1, 1, 3, 2, 4, 0}, findNumOfValidWords(
		[]string{"aaaa", "asas", "able", "ability", "actt", "actor", "access"},
		[]string{"aboveyz", "abrodyz", "abslute", "absoryz", "actresz", "gaswxyz"}))
}

func Test_findNumOfValidWords2(t *testing.T) {
	assert.Equal(t, []int{0, 1, 3, 2, 0}, findNumOfValidWords(
		[]string{"apple", "pleas", "please"},
		[]string{"aelwxyz", "aelpxyz", "aelpsxy", "saelpxy", "xaelpsy"}))
}
