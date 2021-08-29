package task1734

// There is an integer array perm that is a permutation of the first n positive integers, where n is always odd.
// It was encoded into another integer array encoded of length n - 1, such that encoded[i] = perm[i] XOR perm[i + 1].
// For example, if perm = [1,3,2], then encoded = [2,1].
// Given the encoded array, return the original array perm. It is guaranteed that the answer exists and is unique.
//
// Constraints:
//   3 <= n < 105
//   n is odd.
//   x.length == n - 1
//
// Results:
//   Runtime: 156 ms, faster than 100.00% of Go online submissions for Decode XORed Permutation.
//   Memory Usage: 10 MB, less than 37.50% of Go online submissions for Decode XORed Permutation.
func decode(x []int) []int {
	n := len(x) + 1
	result := make([]int, n)

	// Find 1^2^...^n
	xorAll := 0
	for i := 1; i <= n; i++ {
		xorAll ^= i
	}

	// We know that encoded[i] = a[i]^a[i+1], we need to go through pairs:
	// (x[1]=a[1]^a[2], x[3]=a[3]^a[4], ... x[n-1]=a[n-1]^a[n]) and multiply on xorAll
	// a[i] (i != 0) will dissapear
	all := xorAll
	for i := 1; i < n-1; i = i + 2 {
		all ^= x[i]
	}
	result[0] = all

	// When we know first element: we can restore others one by one
	// a[i]^encoded[i] == a[i+1]
	for i := 0; i < n-1; i++ {
		result[i+1] = result[i] ^ x[i]
	}

	return result
}
