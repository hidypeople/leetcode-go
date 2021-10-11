package task0039

// Given an array of distinct integers candidates and a target integer target,
// return a list of all unique combinations of candidates where the chosen
// numbers sum to target. You may return the combinations in any order.
// The same number may be chosen from candidates an unlimited number of times.
// Two combinations are unique if the frequency of at least one of the chosen
// numbers is different.
// It is guaranteed that the number of unique combinations that sum up to target
// is less than 150 combinations for the given input.
//
// Constraints:
//   1 <= candidates.length <= 30
//   1 <= candidates[i] <= 200
//   All elements of candidates are distinct.
//   1 <= target <= 500
//
// Results:
//   Runtime: 0 ms, faster than 100.00% of Go online submissions for Combination Sum.
//   Memory Usage: 3 MB, less than 84.80% of Go online submissions for Combination Sum.
func combinationSum(candidates []int, target int) [][]int {
	return combinationSumPartial(candidates, 0, target)
}

func combinationSumPartial(candidates []int, startFrom int, target int) [][]int {
	result := [][]int{}

	for i := startFrom; i < len(candidates); i++ {
		candidate := candidates[i]
		if candidate < target {
			// Sum of same values
			if target%candidate == 0 {
				res := []int{}
				for l := 1; l <= target/candidate; l++ {
					res = append(res, candidate)
				}
				result = append(result, res)
			}
			// Repeat candidates several times:
			k := 1
			for candidate*k < target {
				// Get subresult for [candidate, candidate, ... combinationSumPartial(...)]
				innerResult := combinationSumPartial(candidates, i+1, target-k*candidate)
				for j := 0; j < len(innerResult); j++ {
					// Append candidate k times for each sub result:
					for l := 1; l <= k; l++ {
						innerResult[j] = append(innerResult[j], candidate)
					}
				}
				result = append(result, innerResult...)
				k++
			}
		} else if candidate > target {
			continue
		} else {
			// Current equals to target
			result = append(result, []int{candidate})
		}
	}

	return result
}
