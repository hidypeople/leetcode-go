package task0125

// Given a string s, determine if it is a palindrome, considering only alphanumeric characters and ignoring cases.
//
// Constraints:
//   1 <= s.length <= 2 * 10^5
//   s consists only of printable ASCII characters.
//
// Results:
//   Runtime: 0 ms, faster than 100.00% of Go online submissions for Valid Palindrome.
//   Memory Usage: 2.7 MB, less than 100.00% of Go online submissions for Valid Palindrome.
func isPalindrome(s string) bool {
	uppercaseDelta := byte('a' - 'A')
	left, right := 0, len(s)-1
	for left < right {
		leftChar, rightChar := s[left], s[right]
		if leftChar >= 'A' && leftChar <= 'Z' {
			leftChar += uppercaseDelta
		}
		if rightChar >= 'A' && rightChar <= 'Z' {
			rightChar += uppercaseDelta
		}
		isLeftLetter := (leftChar >= 'a' && leftChar <= 'z') || (leftChar >= '0' && leftChar <= '9')
		isRightLetter := (rightChar >= 'a' && rightChar <= 'z') || (rightChar >= '0' && rightChar <= '9')
		if isLeftLetter && isRightLetter {
			if leftChar != rightChar {
				return false
			}
			left++
			right--
			continue
		}
		if !isLeftLetter {
			left++
		}
		if !isRightLetter {
			right--
		}
	}
	return true
}
