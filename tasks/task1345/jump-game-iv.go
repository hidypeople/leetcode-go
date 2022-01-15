package task1345

// Given an array of integers arr, you are initially positioned at the first index of the array.
// In one step you can jump from index i to index:
//   i + 1 where: i + 1 < arr.length.
//   i - 1 where: i - 1 >= 0.
//   j where: arr[i] == arr[j] and i != j.
// Return the minimum number of steps to reach the last index of the array.
// Notice that you can not jump outside of the array at any time.
//
// Constraints:
//   1 <= arr.length <= 5 * 104
//   -10^8 <= arr[i] <= 10^8
//
// Results:
//   Runtime: 92 ms, faster than 100.00% of Go online submissions for Jump Game IV.
//   Memory Usage: 13 MB, less than 31.82% of Go online submissions for Jump Game IV.
func minJumps(jumps []int) int {
	n := len(jumps)
	if n == 1 {
		return 0
	}

	directJumps := map[int][]int{}
	for i, v := range jumps {
		directJumps[v] = append(directJumps[v], i)
	}

	visited := make([]bool, n)
	visited[0] = true

	// Use BFS with 1 tweek: if we come to some node go to directJumps and mark them as visited
	jumpsCount := 1
	var nextIndexes []int
	var jump int
	pendingIndexes := []int{0}
	for len(pendingIndexes) > 0 {
		nextIndexes = []int{}
		for _, idx := range pendingIndexes {
			if idx+1 == n-1 {
				return jumpsCount
			}
			if idx+1 < n && !visited[idx+1] {
				// Append next value if it is not visited
				nextIndexes = append(nextIndexes, idx+1)
				visited[idx+1] = true
			}
			if idx-1 >= 0 && !visited[idx-1] {
				// Append prev value if it is not visited
				nextIndexes = append(nextIndexes, idx-1)
				visited[idx-1] = true
			}
			jump = jumps[idx]
			if len(directJumps[jump]) > 1 {
				for _, idx2 := range directJumps[jump] {
					if idx2 == n-1 {
						return jumpsCount
					}
					if !visited[idx2] {
						// Append not visited direct jump items
						nextIndexes = append(nextIndexes, idx2)
						visited[idx2] = true
					}
				}
				// Reset direct jumps to prevent use it again
				directJumps[jump] = nil
			}
		}
		jumpsCount++
		pendingIndexes = nextIndexes
	}
	return jumpsCount
}
