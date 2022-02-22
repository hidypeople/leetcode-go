package task0126

type QueueItem struct {
	word string
	path []string
}

// A transformation sequence from word beginWord to word endWord using a dictionary
// wordList is a sequence of words beginWord -> s1 -> s2 -> ... -> sk such that:
// Every adjacent pair of words differs by a single letter.
// Every si for 1 <= i <= k is in wordList. Note that beginWord does not need to be in wordList.
// sk == endWord
// Given two words, beginWord and endWord, and a dictionary wordList, return all the
// shortest transformation sequences from beginWord to endWord, or an empty list if
// no such sequence exists. Each sequence should be returned as a list of the words
// [beginWord, s1, s2, ..., sk].
//
// Constraints:
//   1 <= beginWord.length <= 5
//   endWord.length == beginWord.length
//   1 <= wordList.length <= 1000
//   wordList[i].length == beginWord.length
//   beginWord, endWord, and wordList[i] consist of lowercase English letters.
//   beginWord != endWord
//   All the words in wordList are unique.
//
// Results:
//   Runtime: 4 ms, faster than 96.55% of Go online submissions for Word Ladder II.
//   Memory Usage: 4.3 MB, less than 27.59% of Go online submissions for Word Ladder II.
func findLadders(beginWord string, endWord string, wordList []string) [][]string {
	if !findString(wordList, endWord) {
		return [][]string{}
	}
	lenWord := len(beginWord)

	// Solve using BFS. Next items are calculated from combinations:
	// "*bc" -> ["abc", "xbc", "zbc"]
	combinations := make(map[string][]string)
	for _, word := range wordList {
		for i := 0; i < lenWord; i++ {
			mask := word[:i] + "*" + word[i+1:]
			combinations[mask] = append(combinations[mask], word)
		}
	}

	queue := make([]QueueItem, 0)
	queue = append(queue, QueueItem{beginWord, []string{beginWord}})
	visited := make(map[string]bool)
	visited[beginWord] = true
	result := make([][]string, 0)

	for len(queue) > 0 {
		lenQueue := len(queue)

		// Go through the whole layer:
		visitedInCurrentLayer := make(map[string]bool)
		for i := 0; i < lenQueue; i++ {
			q := queue[0]
			w := q.word
			p := q.path

			if w == endWord {
				result = append(result, p)
			}

			for j := 0; j < lenWord; j++ {
				mask := w[:j] + "*" + w[j+1:]
				if nextWords, ok := combinations[mask]; ok {
					for _, nextWord := range nextWords {
						if isVisited := visited[nextWord]; !isVisited {
							// make next path
							var nextPath = make([]string, len(p)+1)
							copy(nextPath, p)
							nextPath[len(nextPath)-1] = nextWord

							queue = append(queue, QueueItem{nextWord, nextPath})
							visitedInCurrentLayer[nextWord] = true
						}
					}
				}
			}
			queue = queue[1:]
		}

		// If we met at least one result item -> we don't need to go deeper for longer paths, just return result
		if len(result) > 0 {
			return result
		}

		// We can mark all layer items as visited, because if we'll meet those items
		// again -> that means that will give use worse result:
		// e.g. 'aaa'-> currentLayer=['aab', 'aac'] -> ... -> 'aab'
		//      (this 'aab' is not needed even if we came here through 'aac':
		//      'aaa' -> 'aac' -> ... -> 'aab' because this path larger than 'aaa' -> 'aab')
		for key := range visitedInCurrentLayer {
			visited[key] = true
		}
	}

	return result
}

// findString takes a slice and looks for an element in it. If found it will
// return it's key, otherwise it will return -1 and a bool of false.
func findString(slice []string, val string) bool {
	for _, item := range slice {
		if item == val {
			return true
		}
	}
	return false
}
