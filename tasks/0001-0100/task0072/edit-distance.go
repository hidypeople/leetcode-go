package task0072

// Given two strings word1 and word2, return the minimum number of operations required to convert word1 to word2.
// You have the following three operations permitted on a word:
//   Insert a character
//   Delete a character
//   Replace a character
//
// Constraints:
//   0 <= word1.length, word2.length <= 500
//   word1 and word2 consist of lowercase English letters.
//
// Results:
//   Runtime: 0 ms, faster than 100.00% of Go online submissions for Edit Distance.
//   Memory Usage: 5.9 MB, less than 36.79% of Go online submissions for Edit Distance.
func minDistance(word1 string, word2 string) int {
	// We need to calc "Levenshtein distance": https://en.wikipedia.org/wiki/Levenshtein_distance
	n1, n2 := len(word1), len(word2)

	// Init matrix (e.g. "abc" & "adc"):
	//   1. rows represent the sequence of words from word1: ['', 'a', 'b', 'c']
	//   2. columns represent the sequence of words from word2: ['', 'a', 'd', 'c']
	//   3. position (i,j) contains the number of operations to convert word1[:i] -> word2[:j]
	//      e.g. (1,2) => "a" -> "ad"; (0,3) => "" -> "abc"
	matrix := make([][]int, n1+1)
	for i := 0; i < n1+1; i++ {
		matrix[i] = make([]int, n2+1)
		matrix[i][0] = i
	}
	for i := 1; i < n2+1; i++ {
		matrix[0][i] = i
	}

	for r := 1; r < n1+1; r++ {
		for c := 1; c < n2+1; c++ {
			// for each step we need to find a minimum value between 3 options:
			//    1. left: insert symbol => +1 operation
			//    2. top: delete symbol => +1 operation
			//    3. top-left: replace symbol => +1 operation if symbols not equal, otherwise 0
			left, top, topLeft := matrix[r][c-1], matrix[r-1][c], matrix[r-1][c-1]
			currMin := min(left+1, top+1)
			if word1[r-1] == word2[c-1] {
				currMin = min(currMin, topLeft)
			} else {
				currMin = min(currMin, topLeft+1)
			}
			matrix[r][c] = currMin
		}
	}
	return matrix[n1][n2]
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}
