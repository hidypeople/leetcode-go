package task0135

// There are n children standing in a line. Each child is assigned a rating value given
// in the integer array ratings.
// You are giving candies to these children subjected to the following requirements:
// Each child must have at least one candy.
// Children with a higher rating get more candies than their neighbors.
// Return the minimum number of candies you need to have to distribute the candies to the children.
//
// Constraints:
//   n == ratings.length
//   1 <= n <= 2 * 10^4
//   0 <= ratings[i] <= 2 * 10^4
//
// Results:
//   Runtime: 12 ms, faster than 96.49% of Go online submissions for Candy.
//   Memory Usage: 6.2 MB, less than 85.96% of Go online submissions for Candy.
func candy(ratings []int) int {
	if len(ratings) == 0 {
		return 0
	}
	result := 1
	up, down, peak := 0, 0, 0
	for i := 1; i < len(ratings); i++ {
		if ratings[i-1] < ratings[i] {
			// going up
			up++
			peak = up
			down = 0
			result += 1 + up
		} else if ratings[i-1] == ratings[i] {
			// plain
			peak, up, down = 0, 0, 0
			result += 1
		} else {
			// going down
			up = 0
			down++
			delta := 0
			if peak >= down {
				delta = -1
			}
			result += 1 + down + delta
		}
	}

	return result
}
