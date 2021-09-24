package combination

import "sort"

// Compare 2 sets of combinations
func IsEqualCombinations(arr1 [][]int, arr2 [][]int, isOrderImportant bool) bool {
	for _, a1 := range arr1 {
		found := false
		for _, a2 := range arr2 {
			if len(a1) != len(a2) {
				continue
			}
			isEqual := true
			if isOrderImportant {
				for i := 0; i < len(a1); i++ {
					if a1[i] != a2[i] {
						isEqual = false
					}
				}
				if isEqual {
					found = true
					break
				}
				continue
			}
			copyA1, copyA2 := make([]int, len(a1)), make([]int, len(a1))
			copy(copyA1, a1)
			copy(copyA2, a2)
			sort.Ints(copyA1)
			sort.Ints(copyA2)
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
