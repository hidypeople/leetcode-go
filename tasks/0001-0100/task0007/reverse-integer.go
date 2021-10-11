package task0007

import "strconv"

// Given a signed 32-bit integer x, return x with its digits reversed. If reversing x causes
// the value to go outside the signed 32-bit integer range [-2^31, 2^31 - 1], then return 0.
// Assume the environment does not allow you to store 64-bit integers (signed or unsigned).
//
// Constraints:
//    -2^31 <= x <= 2^31 - 1
//
// Results:
//   Runtime: 0 ms, faster than 100.00% of Go online submissions for Reverse Integer.
//   Memory Usage: 2.2 MB, less than 26.85% of Go online submissions for Reverse Integer.
func reverse(x int) int {
	MaxUint := ^uint32(0)
	MaxInt := int(MaxUint >> 1)
	MinInt := -MaxInt - 1
	minus := x < 0
	if x < 0 {
		x = -x
	}
	s := strconv.Itoa(x)
	if len(s) == 1 {
		return x
	}
	rs := []rune(s)
	for i, j := 0, len(rs)-1; i < j; i, j = i+1, j-1 {
		rs[i], rs[j] = rs[j], rs[i]
	}
	s = string(rs)
	i, e := strconv.Atoi(s)
	if e != nil || i < MinInt || i > MaxInt {
		return 0
	}
	if minus {
		return -i
	}
	return i
}
