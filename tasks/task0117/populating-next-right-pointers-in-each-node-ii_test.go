package task0117

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_connect(t *testing.T) {
	root := &Node{Val: 1}

	root.Left = &Node{Val: 2}
	root.Left.Left = &Node{Val: 4}
	root.Left.Right = &Node{Val: 5}

	root.Right = &Node{Val: 3}
	root.Right.Right = &Node{Val: 7}

	connect(root)
	assert.Nil(t, root.Next)

	assert.Equal(t, root.Left.Next, root.Right)
	assert.Equal(t, root.Left.Left.Next, root.Left.Right)
	assert.Equal(t, root.Left.Right.Next, root.Right.Right)

	assert.Nil(t, root.Right.Next)
	assert.Nil(t, root.Right.Right.Next)
}
