package task0374

// We are playing the Guess Game. The game is as follows:
// I pick a number from 1 to n. You have to guess which number I picked.
// Every time you guess wrong, I will tell you whether the number I picked is higher or lower than your guess.
// You call a pre-defined API int guess(int num), which returns 3 possible results:
//   -1: The number I picked is lower than your guess (i.e. pick < num).
//   1: The number I picked is higher than your guess (i.e. pick > num).
//   0: The number I picked is equal to your guess (i.e. pick == num).
// Return the number that I picked.
//
// Constraints:
//   1 <= n <= 2^31 - 1
//   1 <= pick <= n
//
// Results:
//   Runtime: 0 ms, faster than 100.00% of Go online submissions for Guess Number Higher or Lower.
//   Memory Usage: 2 MB, less than 51.40% of Go online submissions for Guess Number Higher or Lower.
func guessNumberTest(n int, guess func(x int) int) int {
	guessNumber := func(n int) int {
		// Binary search O(lnN)
		left, right := 1, n
		for left < right {
			mid := (left + right) >> 1
			if guess(mid) > 0 {
				left = mid + 1
			} else {
				right = mid
			}
		}
		return left
	}

	return guessNumber(n)
}
