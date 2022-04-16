package task0205

// Two strings s and t are isomorphic if the characters in s can be replaced to get t.
// All occurrences of a character must be replaced with another character while preserving the
// order of characters. No two characters may map to the same character, but a character may map to itself.
//
// Constraints:
//   1 <= s.length <= 5 * 104
//   t.length == s.length
//   s and t consist of any valid ascii character.
//
// Results:
//   Runtime: 0 ms, faster than 100.00% of Go online submissions for Isomorphic Strings.
//   Memory Usage: 2.6 MB, less than 81.12% of Go online submissions for Isomorphic Strings.
func isIsomorphic(s string, t string) bool {
	charsS := make([]byte, 256)
	charsT := make([]byte, 256)
	var chrS byte
	var chrT byte
	var chrSExpected byte
	var chrTExpected byte
	for i := 0; i < len(s); i++ {
		chrS, chrT = s[i], t[i]
		chrSExpected, chrTExpected = charsS[chrS], charsT[chrT]
		if chrSExpected == 0 && chrTExpected == 0 {
			charsS[chrS] = chrT
			charsT[chrT] = chrS
		} else if chrSExpected != chrT || chrTExpected != chrS {
			return false
		}
	}
	return true
}
