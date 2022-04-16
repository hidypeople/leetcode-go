package task0205

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_isIsomorphic(t *testing.T) {
	assert.Equal(t, true, isIsomorphic("egg", "add"))
}

func Test_isIsomorphic2(t *testing.T) {
	assert.Equal(t, false, isIsomorphic("foo", "bar"))
}

func Test_isIsomorphic3(t *testing.T) {
	assert.Equal(t, true, isIsomorphic("paper", "title"))
}

func Test_isIsomorphic4(t *testing.T) {
	assert.Equal(t, false, isIsomorphic("badc", "baba"))
}
