package task0438

// Given two strings s and p, return an array of all the start indices of p's anagrams in s.
// You may return the answer in any order.
// An Anagram is a word or phrase formed by rearranging the letters of a different word or
// phrase, typically using all the original letters exactly once.
//
// Constraints:
//   1 <= s.length, p.length <= 3 * 104
//   s and p consist of lowercase English letters.
//
// Results:
//   Runtime: 4 ms, faster than 96.76% of Go online submissions for Find All Anagrams in a String.
//   Memory Usage: 5 MB, less than 96.76% of Go online submissions for Find All Anagrams in a String.
func findAnagrams(s string, p string) []int {
	nP, nS := len(p), len(s)
	if nP > nS {
		return []int{}
	}
	letters, lettersCurrent := [26]int{}, [26]int{}
	i, result := 0, []int{}
	for i = 0; i < nP; i++ {
		letters[p[i]-'a']++
		lettersCurrent[s[i]-'a']++
	}

	for i = 0; i < nS-nP; i++ {
		if letters == lettersCurrent {
			result = append(result, i)
		}
		// Remove first letter and and next one
		lettersCurrent[s[i]-'a']--
		lettersCurrent[s[i+nP]-'a']++
	}
	if letters == lettersCurrent {
		result = append(result, nS-nP)
	}
	return result
}
