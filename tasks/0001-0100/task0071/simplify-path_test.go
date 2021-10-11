package task0071

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_simplifyPath(t *testing.T) {
	assert.Equal(t, "/home", simplifyPath("/home/"))
}

func Test_simplifyPath2(t *testing.T) {
	assert.Equal(t, "/", simplifyPath("/../"))
}

func Test_simplifyPath3(t *testing.T) {
	assert.Equal(t, "/home/foo", simplifyPath("/home//foo/"))
}

func Test_simplifyPath4(t *testing.T) {
	assert.Equal(t, "/c", simplifyPath("/a/./b/../../c/"))
}
