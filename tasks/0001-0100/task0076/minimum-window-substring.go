package task0076

// Given two strings s and t of lengths m and n respectively, return the minimum window substring of s
// such that every character in t (including duplicates) is included in the window. If there is no
// such substring, return the empty string "".
// The testcases will be generated such that the answer is unique.
// A substring is a contiguous sequence of characters within the string.
//
// Constraints:
//   m == s.length
//   n == t.length
//   1 <= m, n <= 105
//   s and t consist of uppercase and lowercase English letters.
//
// Results:
//   Runtime: 0 ms, faster than 100.00% of Go online submissions for Minimum Window Substring.
//   Memory Usage: 2.9 MB, less than 55.38% of Go online submissions for Minimum Window Substring.
func minWindow(str string, targetStr string) string {
	lenStr, lenTarget := len(str), len(targetStr)
	var targetFrequency, sourceFrequency [256]int
	left, right, finalLeft, finalRight, minWindowWidth, countOfChars := 0, -1, -1, -1, lenStr+1, 0

	for i := 0; i < lenTarget; i++ {
		targetFrequency[targetStr[i]-'a']++
	}

	// First we move right edge forward until we get enough chars,
	// then we move left edge forward until we stop having enough chars
	for left < lenStr {
		if right+1 < lenStr && countOfChars < lenTarget {
			// If we don't have enough chars => continue moving right and collect them
			rightCharIndex := str[right+1] - 'a'
			sourceFrequency[rightCharIndex]++
			if sourceFrequency[rightCharIndex] <= targetFrequency[rightCharIndex] {
				countOfChars++
			}
			right++
		} else {
			// We get enough chars to make window
			if right-left+1 < minWindowWidth && countOfChars == lenTarget {
				minWindowWidth = right - left + 1
				finalLeft = left
				finalRight = right
			}

			// Move left edge forward
			leftCharIndex := str[left] - 'a'
			if sourceFrequency[leftCharIndex] == targetFrequency[leftCharIndex] {
				countOfChars--
			}
			sourceFrequency[leftCharIndex]--
			left++
		}
	}
	if finalLeft != -1 {
		return string(str[finalLeft : finalRight+1])
	}
	return ""
}

// The same idea, but a bit slower (84% cpu)
func minWindow2(str string, targetString string) string {
	if len(str) < len(targetString) {
		return ""
	}

	// Initialize map with pairs {char -> count_of_this_char}
	targetCharsLeft := len(targetString)
	targetChars := make(map[rune]int)
	for _, c := range targetString {
		mapAdd(targetChars, c)
	}

	startResult, endResult := 0, 0
	matchingCharIndexes := make([]int, len(str))
	matchingCharIndexesStart := 0
	matchingCharIndexesEnd := 0
	for i, c := range str {
		charCount, ok := targetChars[c]
		if !ok {
			// Ignore chars that aren't in targetString
			continue
		}
		// If we took the last char: we need to check whether it is last char or not
		if charCount > 0 {
			targetCharsLeft--
		}
		targetChars[c]--
		charCount--

		// keep matching char indexes
		matchingCharIndexes[matchingCharIndexesEnd] = i
		matchingCharIndexesEnd++

		if targetCharsLeft == 0 {
			// The current position contains all chars from targetString!
			// Now we need to extract the smallest substring starting from the beginning
			// of the current start position and ending with current "i":
			newStart, newEnd := -1, i+1
			mIndex := matchingCharIndexesStart
			for mIndex < matchingCharIndexesEnd {
				startIndex := matchingCharIndexes[mIndex]
				newStart = startIndex
				startChar := rune(str[startIndex])
				if targetChars[startChar] == 0 {
					targetCharsLeft = 0
					// If current char is the last possible char break:
					break
				}
				targetChars[startChar]++
				targetCharsLeft++
				mIndex++
			}
			matchingCharIndexesStart = mIndex
			// Update result
			// println(newEnd-newStart, len(result), str[newStart:newEnd])
			if newStart >= 0 && startResult == endResult || newEnd-newStart < endResult-startResult {
				startResult = newStart
				endResult = newEnd
				if endResult-startResult == 1 {
					return str[startResult:endResult]
				}
			}
		}
	}
	return str[startResult:endResult]
}

func mapAdd(hashmap map[rune]int, chr rune) {
	if count, ok := hashmap[chr]; ok {
		hashmap[chr] = count + 1
	} else {
		hashmap[chr] = 1
	}
}
