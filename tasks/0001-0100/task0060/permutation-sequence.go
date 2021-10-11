package task0060

import "strings"

// The set [1, 2, 3, ..., n] contains a total of n! unique permutations.
// By listing and labeling all of the permutations in order, we get the following sequence for n = 3:
// "123" -> "132" -> "213" -> "231" -> "312" -> "321"
// Given n and k, return the kth permutation sequence.
//
// Constraints:
//   1 <= n <= 9
//   1 <= k <= n!
//
// Results:
//   Runtime: 0 ms, faster than 100.00% of Go online submissions for Permutation Sequence.
//   Memory Usage: 2.2 MB, less than 27.27% of Go online submissions for Permutation Sequence.
func getPermutation(n int, k int) string {
	// According to constraints, we have only 9 possible numbers, so it will be easier to store the
	// factorials table instead of calculating factorial each time:
	factorials := map[int]int{
		1: 1,
		2: 2,
		3: 6,
		4: 24,
		5: 120,
		6: 720,
		7: 5040,
		8: 40320,
		9: 362880,
	}
	resultBuilder := new(strings.Builder)

	// Number that have already been used
	numbersDone := make([]bool, n)
	currentK := k
	for i := n - 1; i >= 1; i-- {
		// Calculate symbol position in the current permutation:
		numberOfPermutationsPrev := factorials[i]
		pos := (currentK - 1) / numberOfPermutationsPrev
		currentK = currentK - pos*numberOfPermutationsPrev

		// Find the available number (that wasn't used yet) by current position
		// e.g. available numbers are [1,4,6], pos=1 => we need number "4"
		number := -1
		numberIdx := 0
		for j := 0; j < n && number < 0; j++ {
			if !numbersDone[j] {
				if numberIdx == pos {
					number = j
				}
				numberIdx++
			}
		}
		numbersDone[number] = true
		resultBuilder.WriteByte(byte(number+1) + '0')
	}

	// Process the last number in permutation:
	for j := 0; j < n; j++ {
		if !numbersDone[j] {
			resultBuilder.WriteByte(byte(j+1) + '0')
			break
		}
	}
	return resultBuilder.String()
}
