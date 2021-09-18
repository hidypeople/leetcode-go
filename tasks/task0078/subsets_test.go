package task0078

import (
	"sort"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_subsets(t *testing.T) {
	assert.True(t, isEqualCombinations([][]int{{}, {1}, {2}, {1, 2}, {3}, {1, 3}, {2, 3}, {1, 2, 3}}, subsets([]int{1, 2, 3})))
}

func Test_subsets2(t *testing.T) {
	assert.True(t, isEqualCombinations([][]int{{}, {1}}, subsets([]int{1})))
}

func isEqualCombinations(arr1, arr2 [][]int) bool {
	for _, a1 := range arr1 {
		found := false
		for _, a2 := range arr2 {
			if len(a1) != len(a2) {
				continue
			}
			copyA1, copyA2 := make([]int, len(a1)), make([]int, len(a1))
			copy(copyA1, a1)
			copy(copyA2, a2)
			sort.Ints(copyA1)
			sort.Ints(copyA2)
			isEqual := true
			for i := 0; i < len(a1); i++ {
				if copyA1[i] != copyA2[i] {
					isEqual = false
				}
			}
			if isEqual {
				found = true
				break
			}
		}
		if !found {
			return false
		}
	}
	return true
}
