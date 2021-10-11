package task0068

import (
	"strings"
)

// Given an array of strings words and a width maxWidth, format the text such that each line has exactly
// maxWidth characters and is fully (left and right) justified.
// You should pack your words in a greedy approach; that is, pack as many words as you can in each line.
// Pad extra spaces ' ' when necessary so that each line has exactly maxWidth characters.
// Extra spaces between words should be distributed as evenly as possible. If the number of spaces on a
// line does not divide evenly between words, the empty slots on the left will be assigned more spaces
// than the slots on the right.
// For the last line of text, it should be left-justified and no extra space is inserted between words.
// Note:
//   A word is defined as a character sequence consisting of non-space characters only.
//   Each word's length is guaranteed to be greater than 0 and not exceed maxWidth.
//   The input array words contains at least one word.
//
// Constraints:
//   1 <= words.length <= 300
//   1 <= words[i].length <= 20
//   words[i] consists of only English letters and symbols.
//   1 <= maxWidth <= 100
//   words[i].length <= maxWidth
//
// Results:
//   Runtime: 0 ms, faster than 100.00% of Go online submissions for Text Justification.
//   Memory Usage: 2.4 MB, less than 19.35% of Go online submissions for Text Justification.
func fullJustify(words []string, maxWidth int) []string {
	result := []string{}

	currentWords := []string{}
	currentWordsLen := 0
	for iWord, word := range words {

		// Required space to put current word into the current string:
		spaceRequired := len(word)
		if len(currentWords) > 0 {
			spaceRequired += 1
		}

		// Current space
		currentSpace := currentWordsLen + len(currentWords)

		if currentSpace+spaceRequired-1 <= maxWidth {
			// If fits good - add to current line
			currentWords = append(currentWords, word)
		} else {
			// Otherwise - put to the next line
			s := packWords(currentWords, currentWordsLen, maxWidth, iWord == len(words))
			result = append(result, s)
			currentWords = []string{word}
			currentWordsLen = 0
		}
		currentWordsLen += len(word)
	}
	if len(currentWords) > 0 {
		s := packWords(currentWords, currentWordsLen, maxWidth, true)
		result = append(result, s)
	}

	return result
}

// Put all words into single line with given width. It is guaranteed that words
// can be put into result string at least with one space between them.
func packWords(words []string, wordsLen int, maxWidth int, isLastWord bool) string {
	// Don't need to justify last word in the text
	if isLastWord {
		return strings.Join(words, " ") + strings.Repeat(" ", maxWidth-wordsLen-len(words)+1)
	}
	wordsCount := len(words)
	spacesCount := maxWidth - wordsLen
	result := words[0]
	if wordsCount == 1 {
		// One word: put all spaces at the end of the string
		return result + strings.Repeat(" ", spacesCount)
	} else {
		// Several words: put words justified into the string:
		spaceSlots := wordsCount - 1
		minSpacePerSlot := spacesCount / spaceSlots
		rest := spacesCount - minSpacePerSlot*spaceSlots
		for i := 0; i < spaceSlots; i++ {
			// add spaces
			if rest > 0 {
				result += strings.Repeat(" ", minSpacePerSlot+1)
				rest--
			} else {
				result += strings.Repeat(" ", minSpacePerSlot)
			}
			result += words[i+1]
		}
		return result
	}
}
