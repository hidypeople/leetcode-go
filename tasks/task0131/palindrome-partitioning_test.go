package task0131

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_partition(t *testing.T) {
	assert.Equal(t, [][]string{{"aa", "b"}, {"a", "a", "b"}}, partition("aab"))
}

func Test_partition2(t *testing.T) {
	assert.Equal(t, [][]string{{"a"}}, partition("a"))
}

func Test_partition3(t *testing.T) {
	assert.Equal(t,
		[][]string{{"abba", "b"}, {"a", "bb", "a", "b"}, {"a", "b", "bab"}, {"a", "b", "b", "a", "b"}},
		partition("abbab"))
}

func Test_partition4(t *testing.T) {
	assert.Equal(t,
		[][]string{
			{"cbbbc", "c"},
			{"c", "bbb", "cc"},
			{"c", "bbb", "c", "c"},
			{"c", "bb", "b", "cc"},
			{"c", "bb", "b", "c", "c"},
			{"c", "b", "bb", "cc"},
			{"c", "b", "bb", "c", "c"},
			{"c", "b", "b", "b", "cc"},
			{"c", "b", "b", "b", "c", "c"},
		},
		partition("cbbbcc"))
}
