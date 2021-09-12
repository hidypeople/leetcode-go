package task0066

// You are given a large integer represented as an integer array digits, where
// each digits[i] is the ith digit of the integer. The digits are ordered from most
// significant to least significant in left-to-right order. The large integer does
// not contain any leading 0's.
// Increment the large integer by one and return the resulting array of digits.
//
// Constraints:
//   1 <= digits.length <= 100
//   0 <= digits[i] <= 9
//   digits does not contain any leading 0's.
//
// Results:
//   Runtime: 0 ms, faster than 100.00% of Go online submissions for Plus One.
//   Memory Usage: 2.2 MB, less than 9.80% of Go online submissions for Plus One.
func plusOne(digits []int) []int {
	firstNumber := 0
	digits[len(digits)-1]++
	for j := len(digits) - 1; j >= 0; j-- {
		if digits[j] < 10 {
			break
		} else {
			addCount := digits[j] - 9
			if j > 0 {
				digits[j-1] += addCount
				digits[j] = 0
			} else {
				firstNumber = addCount
				digits[j] = 0
			}
		}
	}
	if firstNumber > 0 {
		digits = append([]int{firstNumber}, digits...)
	}
	return digits
}
