package task0097

// Given strings s1, s2, and s3, find whether s3 is formed by an interleaving of s1 and s2.
// An interleaving of two strings s and t is a configuration where they are divided into non-empty substrings such that:
//   s = s1 + s2 + ... + sn
//   t = t1 + t2 + ... + tm
//   |n - m| <= 1
//   The interleaving is s1 + t1 + s2 + t2 + s3 + t3 + ... or t1 + s1 + t2 + s2 + t3 + s3 + ...
// Note: a + b is the concatenation of strings a and b.
//
// Constraints:
//   0 <= s1.length, s2.length <= 100
//   0 <= s3.length <= 200
//   s1, s2, and s3 consist of lowercase English letters.
//
// Results:
//   Runtime: 0 ms, faster than 100.00% of Go online submissions for Interleaving String.
//   Memory Usage: 2.3 MB, less than 42.86% of Go online submissions for Interleaving String.
func isInterleave(s1 string, s2 string, s3 string) bool {
	n1, n2, n3 := len(s1), len(s2), len(s3)
	if n1+n2 != n3 {
		return false
	}
	// Cache that contains result for s1[i1:] and s2[i2:]
	isInvalid := make([][]bool, n1+1)
	for i := 0; i <= n1; i++ {
		isInvalid[i] = make([]bool, n2+1)
	}

	return dfs(&isInvalid, &s1, &s2, &s3, 0, 0, 0)
}

func dfs(isInvalid *[][]bool, s1, s2, s3 *string, i1, i2, i3 int) bool {
	// Try to hit the cache
	if (*isInvalid)[i1][i2] {
		return false
	}
	if i3 == len(*s3) {
		// if we came to the end of s3 => there is no chars in s1 and s2 because we've checked n1+n2==n3
		return true
	}

	// Check char equality with s3 for s1 and s2 and go deeper if equal
	result :=
		(i1 < len(*s1) && (*s1)[i1] == (*s3)[i3] && dfs(isInvalid, s1, s2, s3, i1+1, i2, i3+1)) ||
			(i2 < len(*s2) && (*s2)[i2] == (*s3)[i3] && dfs(isInvalid, s1, s2, s3, i1, i2+1, i3+1))

	if !result {
		// Update cache value only for falsy results
		(*isInvalid)[i1][i2] = true
	}
	return result
}
