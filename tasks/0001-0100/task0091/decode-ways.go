package task0091

// A message containing letters from A-Z can be encoded into numbers using the following mapping:
//   'A' -> "1"
//   'B' -> "2"
//   ...
//   'Z' -> "26"
// To decode an encoded message, all the digits must be grouped then mapped back into letters using the reverse of the mapping above
// (there may be multiple ways). For example, "11106" can be mapped into:
//   "AAJF" with the grouping (1 1 10 6)
//   "KJF" with the grouping (11 10 6)
// Note that the grouping (1 11 06) is invalid because "06" cannot be mapped into 'F' since "6" is different from "06".
// Given a string s containing only digits, return the number of ways to decode it.
// The answer is guaranteed to fit in a 32-bit integer.
//
// Constraints:
//   1 <= s.length <= 100
//   s contains only digits and may contain leading zero(s).
//
// Results:
//   Runtime: 0 ms, faster than 100.00% of Go online submissions for Decode Ways.
//   Memory Usage: 2.5 MB, less than 15.84% of Go online submissions for Decode Ways.
func numDecodings(s string) int {
	cache := make(map[string]int)
	return numDecodingsRecursive(s, &cache)
}

func numDecodingsRecursive(s string, cache *map[string]int) int {
	s0 := s[0] - '0'
	if s0 == 0 {
		// if current string starts with '0' -> return 0
		return 0
	}
	if res, ok := (*cache)[s]; ok {
		return res
	}
	n := len(s)
	if n == 1 {
		return 1
	}
	s1 := s[1] - '0'
	if n == 2 {
		if s0 > 2 && s1 == 0 {
			// "30", "40",..., "90"
			return 0
		}
		if s1 == 0 || s0 > 2 || (s0 == 2 && s1 > 6) {
			// "10", "20" and "27", "28",..., "99"
			return 1
		}
		return 2
	}
	if s0 == 2 && s1 > 6 || s0 > 2 {
		result := numDecodingsRecursive(s[1:], cache)
		(*cache)[s] = result
		return result
	}
	result := numDecodingsRecursive(s[1:], cache) + numDecodingsRecursive(s[2:], cache)
	(*cache)[s] = result
	return result
}
