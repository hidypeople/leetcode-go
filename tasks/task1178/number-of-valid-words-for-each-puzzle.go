package task1178

// With respect to a given puzzle string, a word is valid if both the following conditions are satisfied:
//    word contains the first letter of puzzle.
//    For each letter in word, that letter is in puzzle.
//       For example, if the puzzle is "abcdefg", then valid words are "faced", "cabbage", and "baggage", while
//       invalid words are "beefed" (does not include 'a') and "based" (includes 's' which is not in the puzzle).
// Return an array answer, where answer[i] is the number of words in the given word list words that is
// valid with respect to the puzzle puzzles[i].
//
// Constraints:
//   1 <= words.length <= 10^5
//   4 <= words[i].length <= 50
//   1 <= puzzles.length <= 10^4
//   puzzles[i].length == 7
//   words[i] and puzzles[i] consist of lowercase English letters.
//   Each puzzles[i] does not contain repeated characters
//
// Results:
//   Runtime: 52 ms, faster than 100.00% of Go online submissions for Number of Valid Words for Each Puzzle.
//   Memory Usage: 15.1 MB, less than 63.64% of Go online submissions for Number of Valid Words for Each Puzzle.
func findNumOfValidWords(words []string, puzzles []string) []int {

	// uniqueWordMasks contains map:
	// mask -> number of matching words for that mask
	uniqueWordMasks := map[int]int{}
	for _, word := range words {
		uniqueWordMasks[mask(word, 0)]++
	}

	result := make([]int, len(puzzles))
	for i, puzzle := range puzzles {
		// mask of first character
		puzzleFirstCharMask := 1 << (puzzle[0] - 'a')
		// mask of puzzle without first character
		puzzleMask := mask(puzzle, 1)

		// Init with the words contains all first puzzle letter
		count := uniqueWordMasks[puzzleFirstCharMask]

		// Start from full mask
		// 1. Find count of words that satisfy to current mask
		// 2. Remove the last '1' bit from current mask on each iteration
		// e.g. our mask 1100100010, first char bit is 1000 ('d' char)
		// check in words map keys: [110010_1_010, 110010_1_000, 110000_1_000, 100000_1_000]
		subsetMask := puzzleMask
		for subsetMask != 0 {
			// Get the words count that matching to "current mask with first bit"
			count += uniqueWordMasks[subsetMask|puzzleFirstCharMask]

			// Remove the last bit in subsetMask
			subsetMask = (subsetMask - 1) & puzzleMask
		}

		result[i] = count
	}
	return result
}

// Make mask: i-th bit equals 1 if corresponding letter exist in the string
func mask(s string, start int) int {
	var result int
	for i := start; i < len(s); i++ {
		result |= 1 << (s[i] - 'a')
	}
	return result
}
