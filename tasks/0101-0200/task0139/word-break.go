package task0139

// Given a string s and a dictionary of strings wordDict, return true if s can be segmented
// into a space-separated sequence of one or more dictionary words.
// Note that the same word in the dictionary may be reused multiple times in the segmentation.
//
// Constraints:
//   1 <= s.length <= 300
//   1 <= wordDict.length <= 1000
//   1 <= wordDict[i].length <= 20
//   s and wordDict[i] consist of only lowercase English letters.
//   All the strings of wordDict are unique.
//
// Results:
//   Runtime: 0 ms, faster than 100.00% of Go online submissions for Word Break.
//   Memory Usage: 2.2 MB, less than 86.60% of Go online submissions for Word Break.
func wordBreak(s string, wordList []string) bool {
	// DP solution

	wordDict := make(map[string]bool, len(wordList))
	for _, word := range wordList {
		wordDict[word] = true
	}

	// dp[i] will contain the result of wordBreak(s[:i], wordList)
	dp := make([]bool, len(s)+1)
	dp[0] = true
	for i := 1; i <= len(s); i++ {
		for j := 0; j < i; j++ {
			if dp[j] && wordDict[s[j:i]] {
				// if s[:j] is good and s[j:i] in the word list => s[:i] is good too
				dp[i] = true
			}
		}
	}
	return dp[len(s)]
}
