package task0006

// The string "PAYPALISHIRING" is written in a zigzag pattern on a given number of rows like this: (you may want to display this pattern in a fixed font for better legibility)
// P   A   H   N
//   A P L S I I G
//    Y   I   R
// And then read line by line: "PAHNAPLSIIGYIR"
//
// Constraints:
//   1 <= s.length <= 1000
//   s consists of English letters (lower-case and upper-case), ',' and '.'.
//   1 <= numRows <= 1000
//
// Results:
//   Runtime: 4 ms, faster than 92.93% of Go online submissions for ZigZag Conversion.
//   Memory Usage: 4 MB, less than 100.00% of Go online submissions for ZigZag Conversion.
func convert(s string, groupNum int) string {
	n := len(s)
	if n == 1 || groupNum == 1 {
		return s
	}
	result := make([]byte, n)
	cycleSize := 2 * (groupNum - 1)
	resultIndex := 0
	for group := 0; group < groupNum; group++ {
		i := 0
		for i < n+cycleSize {
			if group != groupNum-1 && i-group >= 0 && i-group < n {
				result[resultIndex] = s[i-group]
				resultIndex++
			}
			if group > 0 && i+group >= 0 && i+group < n {
				result[resultIndex] = s[i+group]
				resultIndex++
			}
			i += cycleSize
		}
	}
	return string(result)
}
