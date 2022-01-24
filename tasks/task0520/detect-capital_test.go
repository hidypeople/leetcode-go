package task0520

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_detectCapitalUse(t *testing.T) {
	assert.True(t, detectCapitalUse("USA"))
}

func Test_detectCapitalUse2(t *testing.T) {
	assert.False(t, detectCapitalUse("FlAg"))
}

func Test_detectCapitalUse3(t *testing.T) {
	assert.True(t, detectCapitalUse("asd"))
}

func Test_detectCapitalUse4(t *testing.T) {
	assert.True(t, detectCapitalUse("Asdf"))
}
