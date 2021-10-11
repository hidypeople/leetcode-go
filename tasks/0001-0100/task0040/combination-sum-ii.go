package task0040

import "sort"

// Given a collection of candidate numbers (candidates) and a target number (target),
// find all unique combinations in candidates where the candidate numbers sum to target.
// Each number in candidates may only be used once in the combination.
// Note: The solution set must not contain duplicate combinations.
//
// Constraints:
//   1 <= candidates.length <= 100
//   1 <= candidates[i] <= 50
//   1 <= target <= 30
//
// Results:
//   Runtime: 0 ms, faster than 100.00% of Go online submissions for Combination Sum II.
//   Memory Usage: 2.7 MB, less than 80.00% of Go online submissions for Combination Sum II.
func combinationSum2(candidates []int, target int) [][]int {
	sort.Ints(candidates)
	return traverse(candidates, 0, target)
}

func traverse(candidates []int, startIdx int, target int) [][]int {
	result := [][]int{}
	for i := startIdx; i < len(candidates); i++ {
		current := candidates[i]
		if current > target {
			break
		}
		if i > startIdx && candidates[i-1] == current {
			continue
		}

		if current < target {
			out := traverse(candidates, i+1, target-current)
			for k := 0; k < len(out); k++ {
				result = append(result, append(out[k], current))
			}
		} else {
			result = append(result, []int{current})
		}
	}
	return result
}
