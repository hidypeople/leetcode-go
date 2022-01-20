package task0875

import "math"

// Koko loves to eat bananas. There are n piles of bananas, the ith pile has piles[i] bananas.
// The guards have gone and will come back in h hours.
// Koko can decide her bananas-per-hour eating speed of k. Each hour, she chooses some pile of
// bananas and eats k bananas from that pile. If the pile has less than k bananas, she eats all
// of them instead and will not eat any more bananas during this hour.
// Koko likes to eat slowly but still wants to finish eating all the bananas before the guards return.
// Return the minimum integer k such that she can eat all the bananas within h hours.
//
// Constraints:
//   1 <= piles.length <= 10^4
//   piles.length <= h <= 10^9
//   1 <= piles[i] <= 10^9
//
// Results:
//   Runtime: 20 ms, faster than 98.57% of Go online submissions for Koko Eating Bananas.
//   Memory Usage: 6.6 MB, less than 68.57% of Go online submissions for Koko Eating Bananas.
func minEatingSpeed(piles []int, h int) int {
	l, r := 1, 1000000000
	var count int
	var m int
	var canFinish bool
	for l < r {
		m = l + (r-l)>>1
		count = 0
		canFinish = true
		for _, p := range piles {
			count += int(math.Ceil(float64(p) / float64(m)))
			if count > h {
				canFinish = false
				break
			}
		}
		if canFinish {
			r = m
		} else {
			l = m + 1
		}
	}
	return l
}
