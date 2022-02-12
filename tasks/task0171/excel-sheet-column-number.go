package task0171

// Given a string columnTitle that represents the column title as appear in an Excel sheet, return its corresponding column number.
// For example:
//   A -> 1
//   B -> 2
//   C -> 3
//   ...
//   Z -> 26
//   AA -> 27
//   AB -> 28
//
// Constraints:
//   1 <= columnTitle.length <= 7
//   columnTitle consists only of uppercase English letters.
//   columnTitle is in the range ["A", "FXSHRXW"].
//
// Results:
//   Runtime: 0 ms, faster than 100.00% of Go online submissions for Excel Sheet Column Number.
//   Memory Usage: 2.1 MB, less than 68.97% of Go online submissions for Excel Sheet Column Number.
func titleToNumber(columnTitle string) int {
	result, x := 0, 0
	for _, chr := range columnTitle {
		x = int(chr - 'A' + 1)
		result = result*26 + x
	}
	return result
}
