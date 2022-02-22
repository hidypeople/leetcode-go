package task0146

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

const NULL = -1 << 63

func runLruSequence(operations []string, args [][]int) []int {
	results := make([]int, len(operations))
	var lruCache LRUCache
	for i, operation := range operations {
		results[i] = NULL
		switch operation {
		case "LRUCache":
			lruCache = Constructor(args[i][0])
		case "put":
			lruCache.Put(args[i][0], args[i][1])
		case "get":
			results[i] = lruCache.Get(args[i][0])
		}
	}
	return results
}

func TestConstructor(t *testing.T) {
	operations := []string{"LRUCache", "put", "put", "get", "put", "get", "put", "get", "get", "get"}
	args := [][]int{{2}, {1, 1}, {2, 2}, {1}, {3, 3}, {2}, {4, 4}, {1}, {3}, {4}}
	results := runLruSequence(operations, args)
	assert.Equal(t, []int{NULL, NULL, NULL, 1, NULL, -1, NULL, -1, 3, 4}, results)
}

func TestConstructor2(t *testing.T) {
	operations := []string{"LRUCache", "put", "put", "put", "get", "get", "get"}
	args := [][]int{{2}, {1, 100}, {2, 200}, {3, 300}, {1}, {2}, {3}}
	results := runLruSequence(operations, args)
	assert.Equal(t, []int{NULL, NULL, NULL, NULL, -1, 200, 300}, results)
}

func TestConstructor3(t *testing.T) {

	operations := []string{"LRUCache", "put", "get", "put", "get", "get"}
	args := [][]int{{1}, {2, 1}, {2}, {3, 2}, {2}, {3}}
	results := runLruSequence(operations, args)
	assert.Equal(t, []int{NULL, NULL, 1, NULL, -1, 2}, results)
}
