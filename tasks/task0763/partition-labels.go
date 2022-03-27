package task0763

// You are given a string s. We want to partition the string into as many parts as possible so that each letter appears in at most one part.
// Note that the partition is done so that after concatenating all the parts in order, the resultant string should be s.
// Return a list of integers representing the size of these parts.
//
// Constraints:
//   1 <= s.length <= 500
//   s consists of lowercase English letters.
//
// Results:
//   Runtime: 0 ms, faster than 100.00% of Go online submissions for Partition Labels.
//   Memory Usage: 2.1 MB, less than 60.00% of Go online submissions for Partition Labels.
func partitionLabels(s string) []int {

	letterGroups := make([]int, 26)
	letterGroups[s[0]-'a'] = 1
	letterGroupsCount := 1

	for i := 1; i < len(s); i++ {
		letter := s[i]
		groupIdx := letterGroups[letter-'a']
		if groupIdx == 0 {
			// create new group
			letterGroupsCount++
			letterGroups[letter-'a'] = letterGroupsCount
		} else {
			// merge groups within range [groupIdx..]
			for k := 0; k < 26; k++ {
				if groupIdx <= letterGroups[k] {
					letterGroups[k] = groupIdx
				}
			}
			letterGroupsCount = groupIdx
		}
	}

	// Go through the string again and make result according to letterGroups information
	result := []int{}
	prevGroup := 1
	prevGroupIdx := 0
	for i, chr := range s {
		if letterGroups[chr-'a'] == prevGroup {
			continue
		}
		result = append(result, i-prevGroupIdx)
		prevGroup = letterGroups[chr-'a']
		prevGroupIdx = i
	}
	result = append(result, len(s)-prevGroupIdx)

	return result
}
