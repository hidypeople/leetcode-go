package task0119

// Given an integer rowIndex, return the rowIndexth (0-indexed) row of the Pascal's triangle.
//
// Constraints:
//   0 <= rowIndex <= 33
//
// Results:
//   Runtime: 0 ms, faster than 100.00% of Go online submissions for Pascal's Triangle II.
//   Memory Usage: 2 MB, less than 68.00% of Go online submissions for Pascal's Triangle II.
func getRow(n int) []int {
	// O(n) time complexity
	// The idea is that:
	// 1. elements in row n are [C(n,1), C(n,2), ..., C(n, n)], where C(i,j) is binomial coefficent
	// 2. if we know the C(n,k-1), we can calculate C(n,k) = C(n,k-1)*((n-k+1)/k)
	// 3. C(n,k) == C(n,n-k+1) => we need the half of coefficents
	// So we can get all coefficents within range [1..n/2] using recurrent formula
	if n == 0 {
		return []int{1}
	}
	result := make([]int, n+1)
	result[0], result[n] = 1, 1
	var binomialCoefficient int64 = 1
	for i := 1; i <= n/2; i++ {
		binomialCoefficient = (binomialCoefficient * int64(n-i+1)) / int64(i)
		result[i] = int(binomialCoefficient)
		result[n-i] = int(binomialCoefficient)
	}
	return result
}

func getRow2(rowIndex int) []int {
	// Most popular leetcode solution that calculates all pascal triangle values from 0 to rowIndex
	// O(n^2)
	row := make([]int, rowIndex+1)
	for i := 0; i <= rowIndex; i++ {
		row[0] = 1
		row[i] = 1
		for j := i - 1; j > 0; j-- {
			row[j] = row[j-1] + row[j]
		}
	}

	return row
}
