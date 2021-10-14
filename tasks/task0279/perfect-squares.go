package task0279

import "math"

// Given an integer n, return the least number of perfect square numbers that sum to n.
// A perfect square is an integer that is the square of an integer; in other words,
// it is the product of some integer with itself. For example, 1, 4, 9, and 16 are
// perfect squares while 3 and 11 are not.
//
// Constraints:
//   1 <= n <= 10^4
//
// Result:
//   Runtime: 0 ms, faster than 100.00% of Go online submissions for Perfect Squares.
//   Memory Usage: 2.3 MB, less than 98.41% of Go online submissions for Perfect Squares.
func numSquares(n int) int {
	var result = make([]int, 10000)
	if result[n] > 0 {
		return result[n]
	}
	sqrt := sqrtInt(n)
	if n == sqrt*sqrt {
		return 1
	}

	// fill cache with "1" results
	for i := 1; i <= sqrt; i++ {
		result[i*i] = 1
	}

	for i := 2; i <= n; i++ {
		if result[i] > 0 {
			continue
		}
		sqrt = sqrtInt(i)
		min := 4
		for j := 1; j <= sqrt; j++ {
			if min > result[i-j*j]+1 {
				min = result[i-j*j] + 1
			}
		}
		result[i] = min
	}
	return result[n]
}

func sqrtInt(a int) int {
	return int(math.Sqrt(float64(a)))
}
