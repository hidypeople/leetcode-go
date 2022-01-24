package task0520

// We define the usage of capitals in a word to be right when one of the following cases holds:
//   All letters in this word are capitals, like "USA".
//   All letters in this word are not capitals, like "leetcode".
//   Only the first letter in this word is capital, like "Google".
// Given a string word, return true if the usage of capitals in it is right.
//
// Constraints:
//   1 <= word.length <= 100
//   word consists of lowercase and uppercase English letters.
//
// Results:
//   Runtime: 0 ms, faster than 100.00% of Go online submissions for Detect Capital.
//   Memory Usage: 2.2 MB, less than 38.98% of Go online submissions for Detect Capital.
func detectCapitalUse(word string) bool {
	firstCapital := 'A' <= word[0] && word[0] <= 'Z'
	failOnUppercase, failOnLowercase := !firstCapital, false
	for i := 1; i < len(word); i++ {
		isCapital := 'A' <= word[i] && word[i] <= 'Z'
		if (isCapital && failOnUppercase) || (!isCapital && failOnLowercase) {
			return false
		}
		if i == 1 {
			if isCapital {
				failOnLowercase = true
			} else {
				failOnUppercase = true
			}
		}
	}
	return true
}
