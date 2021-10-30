package task1044

// Given a string s, consider all duplicated substrings: (contiguous) substrings of s that
// occur 2 or more times. The occurrences may overlap.
// Return any duplicated substring that has the longest possible length. If s does not
// have a duplicated substring, the answer is "".
//
// Constraints:
//   2 <= s.length <= 3 * 10^4
//   s consists of lowercase English letters.
//
// Results:
//   Runtime: 552 ms, faster than 83.33% of Go online submissions for Longest Duplicate Substring.
//   Memory Usage: 15.7 MB, less than 33.33% of Go online submissions for Longest Duplicate Substring.
func longestDupSubstring(s string) string {
	return longestDupSubstringRabinKarp(s)
}

////////////////////////////////////////////////////////////////////////////////////////////////////////
////////////////////////////////////////////////////////////////////////////////////////////////////////
////////////////////////////////////////////////////////////////////////////////////////////////////////
// Rabin–Karp algorithm
// https://en.wikipedia.org/wiki/Rabin%E2%80%93Karp_algorithm
// Runtime: 552 ms, faster than 83.33% of Go online submissions for Longest Duplicate Substring.
// Memory Usage: 15.7 MB, less than 33.33% of Go online submissions for Longest Duplicate Substring.
////////////////////////////////////////////////////////////////////////////////////////////////////////
////////////////////////////////////////////////////////////////////////////////////////////////////////
////////////////////////////////////////////////////////////////////////////////////////////////////////
func longestDupSubstringRabinKarp(s string) string {
	// Binary search:
	// try the half of string:
	//   if duplicate substrings with half of length found => make it larger, otherwise => make it lesser
	// We need to find the largest string size that have duplicates
	l, r := 0, len(s)
	for l < r {
		m := l + (r-l)/2
		index := testRabinKarp(s, m)
		if index < 0 {
			// if test was unsuccessful: make search string lesser
			r = m
		} else {
			// if test was successful: make search string larger (maybe larger string will be good too)
			l = m + 1
		}
	}
	// l-1 is the LARGEST length of string that CAN pass test
	checkLen := l - 1
	if checkLen <= 0 {
		return ""
	}
	start := testRabinKarp(s, checkLen)
	return s[start : start+checkLen]
}

// Check whether the substring with exact length m is duplicated within the string s
// Calculate hash for each substring with length m and check all substrings with the same hash
// testRabinKarp("banana", 3): check "ban", "ana", "nan", "ana"
func testRabinKarp(s string, m int) int {
	q := (1 << 31) - 1 // large prime
	r := 256           // radix

	// This hashmap contains records:
	// { hashValue: [idx1, idx2, ..., idxN ] } - all indexes that gave the same hash on length m
	mp := make(map[int][]int)

	// Calc hash for mid index and put it into map as key
	// hash === s[0]*r^m-1 + s[1]*r^m-2 + ... s[m-1]
	hash := hashRabinKarp(s, m)
	mp[hash] = []int{0}

	RM := 1 // R^(m-1) mod q
	for i := 1; i <= m-1; i++ {
		RM = (r * RM) % q
	}

	// check range [m..len(s)]
	for i := m; i < len(s); i++ {
		// Remove leading digit, add trailing digit, check for match:
		// H = (s[0]*r^m-1 + ... + s[m-1]) mod q
		// =>
		// H = (s[1]*r^m-1 + ... + s[m]) mod q
		hash = (hash + q - RM*int(s[i-m])%q) % q
		hash = (hash*r + int(s[i])) % q

		startIndex := i - m + 1
		if keys, ok := mp[hash]; ok {
			// if hash was found check possible results
			for _, prev := range keys {
				if compareRabinKarp(s, startIndex, prev, m) {
					return startIndex
				}
			}
		} else {
			mp[hash] = []int{}
		}
		mp[hash] = append(mp[hash], startIndex)
	}

	return -1
}

// Compare s[i1:i1+n] and s[i2:i2+n]
func compareRabinKarp(s string, i1, i2, n int) bool {
	for i := 0; i < n; i++ {
		if s[i1+i] != s[i2+i] {
			return false
		}
	}
	return true
}

// Calc hash for interval s[0:n]
func hashRabinKarp(s string, n int) int {
	q := (1 << 31) - 1 // large prime
	r := 256           // radix

	// hash = (s[0] * r^(n-1) + s[1] * r^(n-2) + ... + s[n-1]) % q
	hash := 0
	for j := 0; j < n; j++ {
		hash = (r*hash + int(s[j])) % q
	}
	return hash
}
