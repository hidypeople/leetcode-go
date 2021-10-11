package task0067

// Given two binary strings a and b, return their sum as a binary string.
//
// Constraints:
//   1 <= a.length, b.length <= 10^4
//   a and b consist only of '0' or '1' characters.
//   Each string does not contain leading zeros except for the zero itself.
//
// Results:
//   Runtime: 0 ms, faster than 100.00% of Go online submissions for Add Binary.
//   Memory Usage: 2.2 MB, less than 92.42% of Go online submissions for Add Binary.
func addBinary(a string, b string) string {
	lenA, lenB := len(a), len(b)
	i, j := lenA-1, lenB-1
	result := make([]byte, max(lenA, lenB)+1)
	k := len(result) - 1
	var addToHigherRank byte = 0
	// Go while there is at least one symbol left
	for i >= 0 || j >= 0 {

		// If both symbols are here
		if i >= 0 && j >= 0 {
			valA, valB := a[i] == '1', b[j] == '1'
			if valA && valB {
				// 1 + 1
				result[k] = '0' + addToHigherRank
				addToHigherRank = 1
			} else if valA || valB {
				// 1 + 0 & 0 + 1
				if addToHigherRank == 1 {
					result[k] = '0'
				} else {
					result[k] = '1'
				}
			} else {
				// 0 + 0
				result[k] = '0' + addToHigherRank
				addToHigherRank = 0
			}
		} else {
			// If only one symbol left

			var val byte
			if i >= 0 {
				val = a[i]
			} else {
				val = b[j]
			}
			if addToHigherRank == 1 {
				if val == '1' {
					result[k] = '0'
				} else {
					result[k] = '1'
					addToHigherRank = 0
				}
			} else {
				result[k] = val
			}
		}
		i--
		j--
		k--
	}
	if addToHigherRank == 1 {
		result[k] = '1'
	} else {
		result = result[1:]
	}
	return string(result)
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
