package task0168

// Given an integer columnNumber, return its corresponding column title as it appears in an Excel sheet.
// A -> 1
// B -> 2
// C -> 3
// ...
// Z -> 26
// AA -> 27
// AB -> 28
//
// Constraints:
//   1 <= columnNumber <= 2^31 - 1
//
// Results:
//   Runtime: 0 ms, faster than 100.00% of Go online submissions for Excel Sheet Column Title.
//   Memory Usage: 1.9 MB, less than 93.27% of Go online submissions for Excel Sheet Column Title.
func convertToTitle(columnNumber int) string {
	result := ""
	var x int
	for columnNumber > 0 {
		x = columnNumber % 26
		if x == 0 {
			x = 26
		}
		result = string('A'+byte(x)-1) + result
		columnNumber = (columnNumber - x) / 26
	}
	return result
}
