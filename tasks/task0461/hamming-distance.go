package task0461

// Given two integers x and y, return the Hamming distance between them.
//
// Constraints:
//   0 <= x, y <= 2^31 - 1
//
// Results:
//   Runtime: 0 ms, faster than 100.00% of Go online submissions for Hamming Distance.
//   Memory Usage: 2 MB, less than 100.00% of Go online submissions for Hamming Distance.
func hammingDistance(x int, y int) int {
	h := x ^ y
	distance := 0
	for h > 0 {
		h = h ^ (h & (-h))
		distance++
	}
	return distance
}
