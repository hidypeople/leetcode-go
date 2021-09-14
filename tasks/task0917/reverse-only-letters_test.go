package task0917

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_reverseOnlyLetters(t *testing.T) {
	assert.Equal(t, "dc-ba", reverseOnlyLetters("ab-cd"))
}

func Test_reverseOnlyLetters2(t *testing.T) {
	assert.Equal(t, "j-Ih-gfE-dCba", reverseOnlyLetters("a-bC-dEf-ghIj"))
}

func Test_reverseOnlyLetters3(t *testing.T) {
	assert.Equal(t, "Qedo1ct-eeLg=ntse-T!", reverseOnlyLetters("Test1ng-Leet=code-Q!"))
}
