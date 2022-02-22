package task0140

// Given a string s and a dictionary of strings wordDict, add spaces in s to construct a
// sentence where each word is a valid dictionary word. Return all such possible sentences in any order.
// Note that the same word in the dictionary may be reused multiple times in the segmentation.
//
// Constraints:
//   1 <= s.length <= 20
//   1 <= wordDict.length <= 1000
//   1 <= wordDict[i].length <= 10
//   s and wordDict[i] consist of only lowercase English letters.
//   All the strings of wordDict are unique.
//
// Results:
//   Runtime: 0 ms, faster than 100.00% of Go online submissions for Word Break II.
//   Memory Usage: 2 MB, less than 59.60% of Go online submissions for Word Break II.
func wordBreak(s string, wordList []string) []string {
	// DP solution
	n, m := len(s), len(wordList)

	// wordDict used for fast search of word in the wordList
	wordDict := make(map[string]int, m)
	for i := 0; i < m; i++ {
		wordDict[wordList[i]] = i + 1
	}

	// DP contains the list of indexes of words from the word list plus 1
	// e.g. "wordXX", ["word", "X", "XX"] =>
	// DP:  [0], [ ], [ ], [ ], [1], [2], [2,3]
	// str: ' '  'w'  'o'  'r'  'd', 'X',  'X'
	// [2,3] means that "X" and "XX" are acceptable if we'll move from the end to beginning of the word
	dp := make([][]int, n+1)
	for i := 0; i < n+1; i++ {
		dp[i] = make([]int, 0)
	}
	dp[0] = []int{0}
	for i := 1; i <= n; i++ {
		for j := 0; j < i; j++ {
			if len(dp[j]) > 0 {
				wordIdx := wordDict[s[j:i]]
				if wordIdx > 0 {
					// if s[:j] is good and s[j:i] in the word list => s[:i] is good too
					dp[i] = append(dp[i], wordIdx)
				}
			}
		}
	}
	if len(dp[n]) == 0 {
		return []string{}
	}

	// Now we need to move from the end of DP and gather the results using DFS
	type StackItem = struct {
		idx int
		s   string
	}
	stack := []StackItem{}
	for _, wordIdx := range dp[n] {
		word := wordList[wordIdx-1]
		stack = append(stack, StackItem{n - len(word), word})
	}
	result := []string{}
	for len(stack) > 0 {
		stackItem := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		if stackItem.idx == 0 {
			result = append(result, stackItem.s)
			continue
		}
		wordIdxs := dp[stackItem.idx]
		for _, wordIdx := range wordIdxs {
			word := wordList[wordIdx-1]
			stack = append(stack, StackItem{stackItem.idx - len(word), word + " " + stackItem.s})
		}
	}
	return result
}
