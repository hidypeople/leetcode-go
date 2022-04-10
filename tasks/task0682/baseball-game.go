package task0682

import "strconv"

// You are keeping score for a baseball game with strange rules. The game consists of several rounds,
// where the scores of past rounds may affect future rounds' scores.
// At the beginning of the game, you start with an empty record. You are given a list of strings ops,
// where ops[i] is the ith operation you must apply to the record and is one of the following:
//   An integer x - Record a new score of x.
//   "+" - Record a new score that is the sum of the previous two scores. It is guaranteed there will always be two previous scores.
//   "D" - Record a new score that is double the previous score. It is guaranteed there will always be a previous score.
//   "C" - Invalidate the previous score, removing it from the record. It is guaranteed there will always be a previous score.
// Return the sum of all the scores on the record.
//
// Constraints:
//   1 <= ops.length <= 1000
//   ops[i] is "C", "D", "+", or a string representing an integer in the range [-3 * 10^4, 3 * 10^4].
//   For operation "+", there will always be at least two previous scores on the record.
//   For operations "C" and "D", there will always be at least one previous score on the record.
//
// Results:
//   Runtime: 0 ms, faster than 100.00% of Go online submissions for Baseball Game.
//   Memory Usage: 2.6 MB, less than 72.22% of Go online submissions for Baseball Game.
func calPoints(ops []string) int {
	result, val := 0, 0
	scores, scoresIdx := make([]int, len(ops)), 0
	for _, op := range ops {
		switch op {
		case "+":
			val = scores[scoresIdx-1] + scores[scoresIdx-2]
			result += val
			scores[scoresIdx] = val
			scoresIdx++
		case "C":
			result -= scores[scoresIdx-1]
			scoresIdx--
		case "D":
			val = 2 * scores[scoresIdx-1]
			result += val
			scores[scoresIdx] = val
			scoresIdx++
		default:
			val, _ = strconv.Atoi(op)
			result += val
			scores[scoresIdx] = val
			scoresIdx++
		}
	}
	return result
}
