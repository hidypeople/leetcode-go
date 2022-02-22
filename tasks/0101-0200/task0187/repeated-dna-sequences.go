package task0187

// The DNA sequence is composed of a series of nucleotides abbreviated as 'A', 'C', 'G', and 'T'.
// For example, "ACGAATTCCG" is a DNA sequence.
// When studying DNA, it is useful to identify repeated sequences within the DNA.
// Given a string s that represents a DNA sequence, return all the 10-letter-long sequences
// (substrings) that occur more than once in a DNA molecule. You may return the answer in any order.
//
// Constraints:
//   1 <= s.length <= 10^5
//   s[i] is either 'A', 'C', 'G', or 'T'.
//
// Results:
//   Runtime: 3 ms, faster than 100.00% of Go online submissions for Repeated DNA Sequences.
//   Memory Usage: 8.2 MB, less than 86.14% of Go online submissions for Repeated DNA Sequences.
func findRepeatedDnaSequences(s string) []string {
	n := len(s)
	if n < 10 {
		return []string{}
	}

	// pack 10-symbol DNA word into 20bit within uint32 variable
	var mask20bit uint32 = 0b00000000_00001111_11111111_11111111
	var packed uint32
	for i := 0; i < 10; i++ {
		chr := s[i]
		packed <<= 2
		switch chr {
		case 'C':
			packed |= 1
		case 'G':
			packed |= 2
		case 'T':
			packed |= 3
		}
	}
	knownDnaWords := map[uint32]int{packed: 1}
	result := []string{}
	for i := 10; i < n; i++ {

		// Append last symbol and erase the first one:
		packed <<= 2
		switch s[i] {
		case 'C':
			packed |= 1
		case 'G':
			packed |= 2
		case 'T':
			packed |= 3
		}
		packed &= mask20bit

		// Check whether we met such symbol or not
		if count := knownDnaWords[packed]; count == 1 {
			result = append(result, s[i-9:i+1])
		}
		knownDnaWords[packed]++
	}

	return result
}
