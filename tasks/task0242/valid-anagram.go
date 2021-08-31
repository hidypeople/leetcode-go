package task0242

// Given two strings s and t, return true if t is an anagram of s, and false otherwise.
//
// Constraints:
// 1 <= s.length, t.length <= 5 * 10^4
// s and t consist of lowercase English letters.
//
// Results:
//   Runtime: 0 ms, faster than 100.00% of Go online submissions for Valid Anagram.
//   Memory Usage: 3 MB, less than 27.41% of Go online submissions for Valid Anagram.
func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}
	hash1, hash2 := make([]int, 26), make([]int, 26)
	for i := 0; i < len(s); i++ {
		hash1[int(s[i]-'a')]++
		hash2[int(t[i]-'a')]++
	}
	for i := 0; i < len(hash1); i++ {
		if hash1[i] != hash2[i] {
			return false
		}
	}
	return true
}
