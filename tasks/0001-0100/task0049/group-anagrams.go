package task0049

// Given an array of strings strs, group the anagrams together. You can return the answer in any order.
// An Anagram is a word or phrase formed by rearranging the letters of a different word or phrase,
// typically using all the original letters exactly once.
//
// Constraints:
// 1 <= strs.length <= 10^4
// 0 <= strs[i].length <= 100
// strs[i] consists of lowercase English letters
//
// Results:
//   Runtime: 16 ms, faster than 98.96% of Go online submissions for Group Anagrams. *Unstable*
//   Memory Usage: 7 MB, less than 99.16% of Go online submissions for Group Anagrams.
func groupAnagrams(strs []string) [][]string {
	intA := int('a')
	hash := make(map[string][]string)

	for _, str := range strs {
		// Calc quick hash key (we use the fact that we have only 26 symbols)
		hashKey := make([]byte, 26)
		for _, chr := range str {
			hashKey[int(chr)-intA]++
		}
		sHash := string(hashKey)

		// Check hash
		if resultArr, ok := hash[sHash]; ok {
			hash[sHash] = append(resultArr, str)
		} else {
			hash[sHash] = []string{str}
		}
	}

	// Move results from hash into results array
	result := make([][]string, len(hash))
	i := 0
	for _, resultArr := range hash {
		result[i] = resultArr
		i++
	}
	return result
}
