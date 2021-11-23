package task0952

var factorsFor100000 []int = []int{
	2, 3, 5, 7, 11, 13, 17, 19, 23, 29, 31, 37, 41, 43, 47, 53, 59, 61, 67, 71, 73, 79, 83, 89, 97, 101,
	103, 107, 109, 113, 127, 131, 137, 139, 149, 151, 157, 163, 167, 173, 179, 181, 191, 193, 197,
	199, 211, 223, 227, 229, 233, 239, 241, 251, 257, 263, 269, 271, 277, 281, 283, 293, 307, 311, 313}

// You are given an integer array of unique positive integers nums. Consider the following graph:
//   There are nums.length nodes, labeled nums[0] to nums[nums.length - 1],
//   There is an undirected edge between nums[i] and nums[j] if nums[i] and nums[j] share
//   a common factor greater than 1.
// Return the size of the largest connected component in the graph.
//
// Constraints:
//   1 <= nums.length <= 2 * 10^4
//   1 <= nums[i] <= 10^5
//   All the values of nums are unique.
//
// Results:
//   Runtime: 172 ms, faster than 100.00% of Go online submissions for Largest Component Size by Common Factor.
//   Memory Usage: 7.4 MB, less than 87.50% of Go online submissions for Largest Component Size by Common Factor.
func largestComponentSize(nums []int) int {
	maxSize := 1

	// "roots" will keep the link to the parent within the group
	// e.g. roots[5] == 3 -> roots[3] == 1 -> roots[1] == 1:
	// (factors with indexes [1,3,5] are in the one group and index 1 is the root index)
	roots := make([]int, len(factorsFor100000))
	// The number of numbers within i-th group
	sizes := make([]int, len(factorsFor100000))
	for i := range roots {
		// A the beginning the root of every number - the number itself
		roots[i] = i
		sizes[i] = 0
	}

	// Get the root of the group of indexes:
	// e.g. i1 -> i2 -> ... -> iN -> iN (iN is the root)
	var getGroupRoot func(int) int
	getGroupRoot = func(i int) int {
		if roots[i] == i {
			return i
		}
		roots[i] = getGroupRoot(roots[i])
		return roots[i]
	}

	// Combine 2 sets:
	//   if they have the same root - do nothing
	//   else set the first to be a parent of the second: root_i <- root_j
	joinGroups := func(i, j int) {
		root_i := getGroupRoot(i)
		root_j := getGroupRoot(j)
		if root_i != root_j {
			// set root_j as parent of root_i
			roots[root_j] = root_i
			sizes[root_i] += sizes[root_j]
			maxSize = max(maxSize, sizes[root_i])
		}
	}

	// It is represents the prime numbers that is greater than primes in factorsFor100000
	largePrimes := map[int]int{}
	for _, num := range nums {
		if num == 1 {
			continue
		}

		last_factor_i := -1
		for i := 0; i < len(factorsFor100000) && factorsFor100000[i] <= num; i++ {
			factor := factorsFor100000[i]
			if num%factor == 0 {
				// factor is divisor of num
				for num%factor == 0 {
					num /= factor
				}
				if last_factor_i >= 0 {
					// Attach current group to the previous one
					joinGroups(last_factor_i, i)
				}
				last_factor_i = i
			}
		}

		if last_factor_i >= 0 {
			// That means that we've already met some small factor
			root_i := getGroupRoot(last_factor_i)
			sizes[root_i]++
			if num > 1 {
				// If num has a large prime
				// x1 * x2 * ... * xN * X (xi - small prime factors, X in range 313 < X <= 100000 and prime number)
				if prime, ok := largePrimes[num]; ok {
					if prime >= 0 {
						// if largePrimes[num] points on existing index in roots
						joinGroups(root_i, prime)
					} else {
						// large prime wasn't connected to roots - get the accumulated counts
						sizes[root_i] += -prime
					}
				}
				// Connect large prime with root_i
				largePrimes[num] = root_i
			}
			maxSize = max(maxSize, sizes[root_i])
		} else {
			// num is large prime: 313 < num <= 100000
			if prime, ok := largePrimes[num]; ok {
				if prime >= 0 {
					// if largePrimes[num] points on existing index in roots
					root_i := getGroupRoot(prime)
					sizes[root_i]++
					maxSize = max(maxSize, sizes[root_i])
				} else {
					// large prime didn't connect to roots: increase count by 1
					largePrimes[num]--
				}
			} else {
				// met large prime first time: set count to 1
				largePrimes[num] = -1
			}
		}
	}

	return maxSize
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
