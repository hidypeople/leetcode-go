package task0048

// You are given an n x n 2D matrix representing an image, rotate the image by 90 degrees (clockwise).
// You have to rotate the image in-place, which means you have to modify the input 2D matrix directly.
// DO NOT allocate another 2D matrix and do the rotation.
//
// Constraints:
//   matrix.length == n
//   matrix[i].length == n
//   1 <= n <= 20
//   -1000 <= matrix[i][j] <= 1000
//
// Results:
//   Runtime: 0 ms, faster than 100.00% of Go online submissions for Rotate Image.
//   Memory Usage: 2.4 MB, less than 20.54% of Go online submissions for Rotate Image.
func rotate(matrix [][]int) {
	n := len(matrix)

	// run from outer cycle into the inner cycles
	for padding := 0; padding < n/2; padding++ {
		currN := n - padding*2       // current cycle square size
		circularN := 4 * (currN - 1) // number of items in current cycle

		// Here we do several cycles with delta shift
		for delta := 0; delta < currN-1; delta++ {
			// Memoize item in current cycle with current delta
			rLast, cLast := circularToMatrix(delta+3*(currN-1), currN)
			valLast := matrix[padding+rLast][padding+cLast]

			// Go reverse clockwise and rotate from next to prev
			for j := delta + 3*(currN-1); j >= currN-1; j -= currN - 1 {
				jNext := (j - (currN - 1)) % circularN
				r1, c1 := circularToMatrix(j, currN)
				r2, c2 := circularToMatrix(jNext, currN)
				matrix[padding+r1][padding+c1] = matrix[padding+r2][padding+c2]
			}

			// Restore item
			matrix[padding][padding+delta] = valLast
		}
	}
}

// From:
// 0 1 2 3
// 11    4
// 10    5
// 9 8 7 6
//
// To: (x,y)
func circularToMatrix(index, n int) (int, int) {
	if index <= n-1 {
		return 0, index // top
	} else if index <= 2*(n-1) {
		return index - (n - 1), n - 1 // right
	} else if index <= 3*(n-1) {
		return n - 1, 3*(n-1) - index // bottom
	} else {
		return 4*(n-1) - index, 0 // left
	}
}
