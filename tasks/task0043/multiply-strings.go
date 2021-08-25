package task0043

import "strings"

// Given two non-negative integers num1 and num2 represented as strings,
// return the product of num1 and num2, also represented as a string.
// Note: You must not use any built-in BigInteger library or convert the inputs to integer directly.
//
// Constraints:
//   1 <= num1.length, num2.length <= 200
//   num1 and num2 consist of digits only.
//   Both num1 and num2 do not contain any leading zero, except the number 0 itself.
//
// Results:
//   Runtime: 0 ms, faster than 100.00% of Go online submissions for Multiply Strings.
//   Memory Usage: 5.9 MB, less than 31.49% of Go online submissions for Multiply Strings.
func multiply(num1 string, num2 string) string {
	if num1 == "0" || num2 == "0" {
		return "0"
	}
	n1 := []byte(num1)
	n2 := []byte(num2)
	// Reverse both string for easier work with them (complexity: O(N))
	reverseString(n1)
	reverseString(n2)
	result := make([]int, len(n1)+len(n2)+1)
	pow := 0
	// Run over first number digits lower->higher digit rank
	for i := 0; i < len(n1); i++ {
		val := n1[i] - '0'
		// Multiply current digit to whole second number
		multResult := multiplyOne(n2, int(val), pow)

		// Add result
		var reminder int = 0
		for j := 0; j < len(multResult); j++ {
			prev := result[j]
			result[j] = (prev + multResult[j] + reminder) % 10
			reminder = (prev + multResult[j] + reminder) / 10
		}
		result[len(multResult)] += reminder
		pow++
	}
	// Build result
	builder := new(strings.Builder)
	isZero := true
	for i := len(result) - 1; i >= 0; i-- {
		if isZero && result[i] == 0 {
			continue
		}
		if result[i] != 0 {
			isZero = false
		}
		builder.WriteByte(byte(result[i] + '0'))
	}
	return builder.String()
}

// Multiply single digit number to large number
func multiplyOne(num []byte, val int, pow int) []int {
	if val == 0 {
		return make([]int, pow)
	}
	n := len(num)
	result := make([]int, n+pow+1)
	reminder := 0
	resultIndex := pow
	for i := pow; i < n+pow; i++ {
		mult := val*int(num[i-pow]-'0') + reminder
		result[resultIndex] = mult % 10
		reminder = mult / 10
		resultIndex++
	}
	result[n+pow] = reminder
	for i := 0; i < pow; i++ {
		result[i] = 0
	}
	return result
}

func reverseString(bytes []byte) {
	for i, j := 0, len(bytes)-1; i < j; i, j = i+1, j-1 {
		bytes[i], bytes[j] = bytes[j], bytes[i]
	}
}
