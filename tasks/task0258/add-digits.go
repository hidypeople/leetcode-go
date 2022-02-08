package task0258

// Given an integer num, repeatedly add all its digits until the result has only one digit, and return it.
//
// Constraints:
//   0 <= num <= 2^31 - 1
//
// Results:
//   Runtime: 0 ms, faster than 100.00% of Go online submissions for Add Digits.
//   Memory Usage: 2.1 MB, less than 75.00% of Go online submissions for Add Digits.
func addDigits(num int) int {
	for num >= 10 {
		curr, sum := num, 0
		for curr > 0 {
			sum += curr % 10
			curr = curr / 10
		}
		num = sum
	}
	return num
}
