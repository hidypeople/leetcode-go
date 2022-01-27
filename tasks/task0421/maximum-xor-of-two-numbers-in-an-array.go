package task0421

// Given an integer array nums, return the maximum result of nums[i] XOR nums[j], where 0 <= i <= j < n.
//
// Constraints:
//   1 <= nums.length <= 2 * 10^5
//   0 <= nums[i] <= 2^31 - 1
//
// Results:
//   Runtime: 128 ms, faster than 100.00% of Go online submissions for Maximum XOR of Two Numbers in an Array.
//   Memory Usage: 13.1 MB, less than 33.33% of Go online submissions for Maximum XOR of Two Numbers in an Array.
func findMaximumXOR(nums []int) int {
	ans, root := 0, &Trie{}
	var current *Trie
	var bit int
	for _, num := range nums {
		current = root
		for i := 31; i >= 0; i-- {
			bit = (num >> i) & 1
			if current.child[bit] == nil {
				current.child[bit] = &Trie{}
			}
			current = current.child[bit]
		}
	}
	var sum int
	for _, num := range nums {
		sum = 0
		current = root
		for i := 31; i >= 0; i-- {
			bit = (num >> i) & 1
			if current.child[(bit^1)] != nil {
				sum += 1 << i
				current = current.child[(bit ^ 1)]
			} else {
				current = current.child[bit]
			}
		}
		ans = max(sum, ans)
	}
	return ans
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

type Trie struct {
	child [2]*Trie
}
