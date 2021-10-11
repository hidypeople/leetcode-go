package task0030

// You are given a string s and an array of strings words of the same length.
// Return all starting indices of substring(s) in s that is a concatenation of each word in words exactly once, in any order, and without any intervening characters.
// You can return the answer in any order.
//
// Constraints:
//   1 <= s.length <= 10^4
//   s consists of lower-case English letters.
//   1 <= words.length <= 5000
//   1 <= words[i].length <= 30
//   words[i] consists of lower-case English letters
//
// Results:
//   Runtime: 4 ms, faster than 100.00% of Go online submissions for Substring with Concatenation of All Words.
//   Memory Usage: 4 MB, less than 84.38% of Go online submissions for Substring with Concatenation of All Words.
func findSubstring(s string, words []string) []int {
	wordLen := len(words[0])
	wordToCount := make(map[string]int16)
	for _, word := range words {
		wordToCount[word]++
	}

	// This func revert wordToCount values for all items in buffer
	// until we meet targetWord, buffer items cut offs (including first targetWord):
	//                                  buffer=["a", "b", "c", "d"], wordToCount=[a:0, b:0, c:1, d:2]
	// revertWordBuffer(buffer, "c") => buffer=["d"],                wordToCount=[a:1, b:1, c:2, d:2]
	revertWordBuffer := func(buffer []string, targetWord string) []string {
		if len(buffer) == 0 {
			return buffer
		}
		foundIndex := len(buffer)
		for i, word := range buffer {
			wordToCount[word]++
			if word == targetWord {
				foundIndex = i
				break
			}
		}
		if foundIndex == len(buffer) {
			return []string{}
		}
		return buffer[foundIndex+1:]
	}

	result := []int{}

	// start from 0 -> start from 1 -> ... -> start from wordLen-1
	for delta := 0; delta < wordLen; delta++ {

		startIndex := delta
		wordBuffer := []string{}
		for i := delta; i+wordLen <= len(s); i += wordLen {
			word := s[i : i+wordLen]
			if availableCount, ok := wordToCount[word]; ok {
				if availableCount <= 0 {
					// If there is no available items
					oldBufferLen := len(wordBuffer)
					wordBuffer = revertWordBuffer(wordBuffer, word)
					startIndex += wordLen * (oldBufferLen - len(wordBuffer))
					wordBuffer = append(wordBuffer, word)
					wordToCount[word]--
				} else {
					// If word is okay and avaialable append to buffer:
					wordBuffer = append(wordBuffer, word)
					wordToCount[word]--
					if len(wordBuffer) == len(words) {
						result = append(result, startIndex)
						startIndex += wordLen
						wordToCount[wordBuffer[0]]++
						wordBuffer = wordBuffer[1:]
					}
				}
			} else {
				// Reset start index
				startIndex = i + wordLen
				// Reset word buffer
				wordBuffer = revertWordBuffer(wordBuffer, "")
				continue
			}
		}
		wordBuffer = revertWordBuffer(wordBuffer, "")
	}

	return result
}
