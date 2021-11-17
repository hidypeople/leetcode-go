package task0127

type QueueItem struct {
	word string
	len  int
}

// A transformation sequence from word beginWord to word endWord using a dictionary wordList is
// a sequence of words beginWord -> s1 -> s2 -> ... -> sk such that:
//   Every adjacent pair of words differs by a single letter.
//   Every si for 1 <= i <= k is in wordList. Note that beginWord does not need to be in wordList.
//   sk == endWord
// Given two words, beginWord and endWord, and a dictionary wordList, return the number of
// words in the shortest transformation sequence from beginWord to endWord, or 0 if no such sequence exists.
//
// Constraints:
//   1 <= beginWord.length <= 10
//   endWord.length == beginWord.length
//   1 <= wordList.length <= 5000
//   wordList[i].length == beginWord.length
//   beginWord, endWord, and wordList[i] consist of lowercase English letters.
//   beginWord != endWord
//   All the words in wordList are unique.
//
// Results:
//   Runtime: 40 ms, faster than 95.97% of Go online submissions for Word Ladder.
//   Memory Usage: 9.5 MB, less than 23.49% of Go online submissions for Word Ladder.
func ladderLength(beginWord string, endWord string, wordList []string) int {
	if beginWord == endWord {
		return 1
	}
	// Ensure that end word exist in the word list
	endWordFound := false
	for _, word := range wordList {
		if word == endWord {
			endWordFound = true
		}
	}
	if !endWordFound {
		return 0
	}

	// Solve using BFS. Next items are calculated from combinations:
	// "*bc" -> ["abc", "xbc", "zbc"]
	lenWord := len(beginWord)
	combinations := make(map[string][]string)
	for _, word := range wordList {
		for i := 0; i < lenWord; i++ {
			mask := word[:i] + "*" + word[i+1:]
			combinations[mask] = append(combinations[mask], word)
		}
	}

	// BFS
	visited := map[string]bool{beginWord: true}
	queue := []QueueItem{{beginWord, 1}}
	for len(queue) > 0 {

		// run through current level
		nextQueue := []QueueItem{}
		for _, queueItem := range queue {
			word := queueItem.word
			length := queueItem.len
			for i := 0; i < lenWord; i++ {
				mask := word[:i] + "*" + word[i+1:]
				nextWords := combinations[mask]
				for _, nextWord := range nextWords {
					// If we reach the end
					if nextWord == endWord {
						return length + 1
					}
					if visited[nextWord] {
						continue
					}
					nextQueue = append(nextQueue, QueueItem{nextWord, length + 1})
					visited[nextWord] = true
				}
			}
		}
		queue = nextQueue
	}
	return 0
}
