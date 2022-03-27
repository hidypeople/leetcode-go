package task0763

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_partitionLabels(t *testing.T) {
	assert.Equal(t, []int{9, 7, 8}, partitionLabels("ababcbacadefegdehijhklij"))
}

func Test_partitionLabels2(t *testing.T) {
	assert.Equal(t, []int{10}, partitionLabels("eccbbbbdec"))
}
