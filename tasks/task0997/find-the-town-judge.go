package task0997

// In a town, there are n people labeled from 1 to n. There is a rumor that one of these people is secretly the town judge.
// If the town judge exists, then:
//   The town judge trusts nobody.
//   Everybody (except for the town judge) trusts the town judge.
//   There is exactly one person that satisfies properties 1 and 2.
//   You are given an array trust where trust[i] = [ai, bi] representing that the person labeled ai trusts the person labeled bi.
// Return the label of the town judge if the town judge exists and can be identified, or return -1 otherwise.
//
// Constraints:
//   1 <= n <= 1000
//   0 <= trust.length <= 104
//   trust[i].length == 2
//   All the pairs of trust are unique.
//   ai != bi
//   1 <= ai, bi <= n
//
// Results:
//   Runtime: 96 ms, faster than 98.80% of Go online submissions for Find the Town Judge.
//   Memory Usage: 7.9 MB, less than 30.12% of Go online submissions for Find the Town Judge.
func findJudge(n int, trust [][]int) int {
	trustedCount := make([]int, n)
	for _, x := range trust {
		trustedCount[x[0]-1]--
		trustedCount[x[1]-1]++
	}
	for i, x := range trustedCount {
		if x == n-1 {
			return i + 1
		}
	}
	return -1
}
