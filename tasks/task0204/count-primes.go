package task0204

// Given an integer n, return the number of prime numbers that are strictly less than n.
//
// Constraints:
//   0 <= n <= 5 * 10^6
//
// Results:
//   Runtime: 34 ms, faster than 98.81% of Go online submissions for Count Primes.
//   Memory Usage: 13.1 MB, less than 44.05% of Go online submissions for Count Primes.
func countPrimes(n int) int {
	if n < 3 {
		return 0
	}

	nonPrimes := make([]bool, n)
	var i int
	var j int

	// Maximum of prime number count within range [2..n]
	result := n / 2
	// Start from 3-rd item and skip everything that can be divided by 2
	for i = 3; i*i < n; i += 2 {
		if nonPrimes[i] {
			continue
		}
		for j = i * i; j < n; j += 2 * i {
			if !nonPrimes[j] {
				// if this number previously has been considered as prime - reduce result count
				result--
			}
			nonPrimes[j] = true
		}
	}
	return result
}
