package task0899

import "sort"

// You are given a string s and an integer k. You can choose one of the first k letters of s and append it at the end of the string..
// Return the lexicographically smallest string you could have after applying the mentioned step any number of moves.
//
// Constraints:
//   1 <= k <= s.length <= 1000
//   s consist of lowercase English letters.
//
// Results:
//   Runtime: 0 ms, faster than 100.00% of Go online submissions for Orderly Queue.
//   Memory Usage: 2.2 MB, less than 100.00% of Go online submissions for Orderly Queue.
func orderlyQueue(s string, k int) string {
	sorted := []byte(s)
	sort.Slice(sorted, func(i, j int) bool {
		return sorted[i] < sorted[j]
	})
	// If we have k >= 2 - that means that we can swap any 2 items => we can sort this string
	// lexicographically using the basic operation many times
	if k >= 2 {
		return string(sorted)
	}

	return rotate(s, sorted[0])
}

// we need to find all rotated strings starting from the minimum character
// and lexicographically smaller than others
func rotate(str string, targetSymbol byte) string {
	targets := []string{}
	for i := 0; i < len(str); i++ {
		if str[i] == targetSymbol {
			targets = append(targets, str[:0]+str[i:]+str[0:i])
		}
	}
	if len(targets) == 0 {
		return str
	}
	sort.Slice(targets, func(i, j int) bool {
		return compare(targets[i], targets[j]) < 0
	})
	return targets[0]
}

func compare(s1, s2 string) int {
	for i := 0; i < len(s1); i++ {
		if s1[i] < s2[i] {
			return -1
		} else if s1[i] > s2[i] {
			return 1
		}
	}
	return 0
}
