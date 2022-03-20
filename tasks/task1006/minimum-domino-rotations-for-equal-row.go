package task1006

import "math"

// In a row of dominoes, tops[i] and bottoms[i] represent the top and bottom halves of the ith domino.
// (A domino is a tile with two numbers from 1 to 6 - one on each half of the tile.)
// We may rotate the ith domino, so that tops[i] and bottoms[i] swap values.
// Return the minimum number of rotations so that all the values in tops are the same,
// or all the values in bottoms are the same.
// If it cannot be done, return -1.
//
// Constraints:
//   2 <= tops.length <= 2 * 10^4
//   bottoms.length == tops.length
//   1 <= tops[i], bottoms[i] <= 6
//
// Results:
//   Runtime: 104 ms, faster than 100.00% of Go online submissions for Minimum Domino Rotations For Equal Row.
//   Memory Usage: 7.3 MB, less than 78.85% of Go online submissions for Minimum Domino Rotations For Equal Row.
func minDominoRotations(tops []int, bottoms []int) int {
	n := len(tops)
	result := math.MaxInt32
	var top int
	var bottom int
	for curr := 1; curr <= 6; curr++ {
		// count of flips for top and bottom domino parts (-1 - means impossible for "curr")
		flipsTop, flipsBottom := 0, 0
		for i := 0; i < n; i++ {
			top, bottom = tops[i], bottoms[i]

			if top != curr && bottom != curr {
				// both top and bottom isn't equal to curr
				flipsTop = -1
				flipsBottom = -1
				break
			}

			if flipsTop != -1 && top != curr {
				// top is not equals to curr domino value => try to flip
				if bottom != curr {
					// another part doesn't fit too
					flipsTop = -1
					if flipsBottom == -1 {
						break
					}
				} else {
					flipsTop++
				}
			}
			if flipsBottom != -1 && bottom != curr {
				// bottom is not equals to curr domino value => try to flip
				if top != curr {
					// another part doesn't fit too
					flipsBottom = -1
					if flipsTop == -1 {
						break
					}
				} else {
					flipsBottom++
				}
			}
		}
		if flipsTop != -1 {
			result = min(result, flipsTop)
		}
		if flipsBottom != -1 {
			result = min(result, flipsBottom)
		}
	}
	if result == math.MaxInt32 {
		return -1
	}
	return result
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}
