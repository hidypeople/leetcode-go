package task0077

// Given two integers n and k, return all possible combinations of k numbers out of the range [1, n].
// You may return the answer in any order.
//
// Constraints:
//   1 <= n <= 20
//   1 <= k <= n
//
// Results:
//   Runtime: 8 ms, faster than 91.72% of Go online submissions for Combinations.
//   Memory Usage: 7.3 MB, less than 53.31% of Go online submissions for Combinations.
func combine(n int, k int) [][]int {
	if k == 1 {
		res := [][]int{}
		for i := 1; i <= n; i++ {
			res = append(res, []int{i})
		}
		return res
	}
	return combineNumbers(0, n, k, k)
}

func combineNumbers(idx, n, kOriginal, k int) [][]int {
	result := [][]int{}
	for i := idx; i < n-k+1; i++ {
		currentNumber := i + 1
		var res [][]int
		if k == 2 {
			res = make([][]int, n-i-1)
			for j := i + 1; j < n; j++ {
				a := make([]int, kOriginal)
				a[kOriginal-1] = j + 1
				res[j-i-1] = a
			}
		} else {
			res = combineNumbers(i+1, n, kOriginal, k-1)
		}
		for i := range res {
			res[i][kOriginal-k] = currentNumber
			result = append(result, res[i])
		}
	}
	return result
}
