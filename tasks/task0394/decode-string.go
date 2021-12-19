package task0394

// Given an encoded string, return its decoded string.
// The encoding rule is: k[encoded_string], where the encoded_string inside
// the square brackets is being repeated exactly k times. Note that k is guaranteed to be a positive integer.
// You may assume that the input string is always valid; No extra white spaces, square brackets are well-formed, etc.
// Furthermore, you may assume that the original data does not contain any digits and that
// digits are only for those repeat numbers, k. For example, there won't be input like 3a or 2[4].
//
// Constraints:
//   1 <= s.length <= 30
//   s consists of lowercase English letters, digits, and square brackets '[]'.
//   s is guaranteed to be a valid input.
//   All the integers in s are in the range [1, 300].
//
// Results:
//   Runtime: 0 ms, faster than 100.00% of Go online submissions for Decode String.
//   Memory Usage: 2.3 MB, less than 25.47% of Go online submissions for Decode String.
func decodeString(s string) string {
	result, _ := decodeStringPart(0, s)
	return result
}

func decodeStringPart(idx int, s string) (string, int) {
	result := ""
	i := idx
	for i < len(s) {
		if s[i] >= 'a' && s[i] <= 'z' {
			// letter => add to result
			result = result + string(s[i])
			i++
		} else if s[i] >= '0' && s[i] <= '9' {
			// number => get the number and start substring
			repeats := 0
			for s[i] != '[' {
				repeats = repeats*10 + int(s[i]-'0')
				i++
			}
			subStringResult, newIdx := decodeStringPart(i+1, s)
			i = newIdx
			for j := 0; j < repeats; j++ {
				result += subStringResult
			}
		} else if s[i] == ']' {
			return result, i + 1
		}
	}
	return result, i + 1
}
