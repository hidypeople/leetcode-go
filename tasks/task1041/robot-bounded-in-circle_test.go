package task1041

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_isRobotBounded(t *testing.T) {
	assert.True(t, isRobotBounded("GGLLGG"))
}

func Test_isRobotBounded2(t *testing.T) {
	assert.False(t, isRobotBounded("GG"))
}

func Test_isRobotBounded3(t *testing.T) {
	assert.True(t, isRobotBounded("GL"))
}

func Test_isRobotBounded4(t *testing.T) {
	assert.True(t, isRobotBounded("GLRLGLLGLGRGLGL"))
}
