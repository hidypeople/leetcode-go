package task0567

// Given two strings subStr and str, return true if str contains a permutation of s1, or false otherwise.
// In other words, return true if one of subStr's permutations is the substring of str.
//
// Constraints:
//   1 <= subStr.length, str.length <= 10^4
//   subStr and str consist of lowercase English letters.
//
// Results:
//   Runtime: 0 ms, faster than 100.00% of Go online submissions for Permutation in String.
//   Memory Usage: 2.5 MB, less than 100.00% of Go online submissions for Permutation in String.
func checkInclusion(subStr string, str string) bool {
	nSub, n := len(subStr), len(str)
	if n < nSub {
		return false
	}
	strLetters, subStrLetters := [26]int{}, [26]int{}
	for i, chr := range subStr {
		strLetters[str[i]-'a']++
		subStrLetters[chr-'a']++
	}
	for i := nSub; i < n; i++ {
		// check existing letters
		if strLetters == subStrLetters {
			return true
		}
		strLetters[str[i-nSub]-'a']--
		strLetters[str[i]-'a']++
	}
	return strLetters == subStrLetters
}
