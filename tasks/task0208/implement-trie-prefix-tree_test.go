package task0208

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

type Bool interface{}

var trueVal Bool = true
var falseVal Bool = false

func TestTrie_Insert(t *testing.T) {
	RunOperations(t,
		[]string{"Trie", "insert", "search", "search", "startsWith", "insert", "search"},
		[][]string{{}, {"apple"}, {"apple"}, {"app"}, {"app"}, {"app"}, {"app"}},
		[]Bool{nil, nil, trueVal, falseVal, trueVal, nil, trueVal})
}

func TestTrie_Insert2(t *testing.T) {
	RunOperations(t,
		[]string{
			"Trie",
			"insert", "insert", "insert",
			"insert", "insert", "insert",
			"search", "search", "search", "search",
			"startsWith", "startsWith", "startsWith", "startsWith",
		},
		[][]string{
			{},
			{"aaaaa"}, {"aaaab"}, {"aaaac"},
			{"aaaa"}, {"aaab"}, {"aaac"},
			{"aaaaa"}, {"aaaaaa"}, {"aaaa"}, {"a"},
			{"aaaaa"}, {"aaaaaa"}, {"aaaa"}, {"a"},
		},
		[]Bool{
			nil,
			nil, nil, nil,
			nil, nil, nil,
			trueVal, falseVal, trueVal, falseVal,
			trueVal, falseVal, trueVal, trueVal,
		})
}

func TestTrie_Insert3(t *testing.T) {
	RunOperations(t,
		[]string{"Trie", "insert", "insert", "insert", "search", "search", "search", "startsWith"},
		[][]string{{}, {"aaa"}, {"aaaaa"}, {"aa"}, {"aa"}, {"aaa"}, {"aaaa"}, {"aaaaa"}, {"a"}, {"aaaa"}},
		[]Bool{nil, nil, nil, nil, trueVal, trueVal, falseVal, trueVal, falseVal, trueVal})
}

func RunOperations(t *testing.T, operations []string, params [][]string, expected []Bool) {
	var result Trie
	for opIdx, op := range operations {
		switch op {
		case "Trie":
			result = Constructor()
		case "insert":
			result.Insert(params[opIdx][0])
		case "search":
			if expected[opIdx] != result.Search(params[opIdx][0]) {
				fmt.Printf("Search(%v)\n", params[opIdx][0])
			}
			assert.Equal(t, expected[opIdx], result.Search(params[opIdx][0]))
		case "startsWith":
			if expected[opIdx] != result.StartsWith(params[opIdx][0]) {
				fmt.Printf("StartsWith(%v)\n", params[opIdx][0])
			}
			assert.Equal(t, expected[opIdx], result.StartsWith(params[opIdx][0]))
		}
	}
}
