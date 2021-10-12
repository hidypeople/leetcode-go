package task0374

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func makeGuessFn(number int) func(int) int {
	return func(x int) int {
		if x < number {
			return 1
		} else if x == number {
			return 0
		} else {
			return -1
		}
	}
}

func Test_guessNumber(t *testing.T) {
	n := 20
	for i := 1; i <= n; i++ {
		assert.Equal(t, i, guessNumberTest(n, makeGuessFn(i)))
	}
}
