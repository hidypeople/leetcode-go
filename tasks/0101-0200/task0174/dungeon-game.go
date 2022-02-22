package task0174

// The demons had captured the princess and imprisoned her in the bottom-right corner of a dungeon.
// The dungeon consists of m x n rooms laid out in a 2D grid. Our valiant knight was initially positioned
// in the top-left room and must fight his way through dungeon to rescue the princess.
// The knight has an initial health point represented by a positive integer. If at any point his health
// point drops to 0 or below, he dies immediately.
// Some of the rooms are guarded by demons (represented by negative integers), so the knight loses health
// upon entering these rooms; other rooms are either empty (represented as 0) or contain magic orbs that
// increase the knight's health (represented by positive integers).
// To reach the princess as quickly as possible, the knight decides to move only rightward or downward in each step.
// Return the knight's minimum initial health so that he can rescue the princess.
// Note that any room can contain threats or power-ups, even the first room the knight enters and the
// bottom-right room where the princess is imprisoned.
//
// Constraints:
//   m == dungeon.length
//   n == dungeon[i].length
//   1 <= m, n <= 200
//   -1000 <= dungeon[i][j] <= 1000
//
// Results:
//   Runtime: 4 ms, faster than 90.63% of Go online submissions for Dungeon Game.
//   Memory Usage: 3.7 MB, less than 34.38% of Go online submissions for Dungeon Game.
func calculateMinimumHP(dungeon [][]int) int {
	m, n := len(dungeon), len(dungeon[0])

	dp := make([][]int, m+1)
	for i := 0; i < m+1; i++ {
		dp[i] = make([]int, n+1)
	}
	// We know that we should have 1HP at the end of the maze,
	// so we need to make a conditions that will allow us this:
	dp[m][n-1] = 1
	dp[m-1][n] = 1

	for r := m - 1; r >= 0; r-- {
		for c := n - 1; c >= 0; c-- {
			// The minimum cost that is required to come to this room from right and bottom
			// => plus apply current room value
			bottom, right := dp[r+1][c], dp[r][c+1]
			var health int
			if bottom > 0 && right > 0 {
				health = min(bottom, right) - dungeon[r][c]
			} else if bottom == 0 {
				health = right - dungeon[r][c]
			} else {
				health = bottom - dungeon[r][c]
			}

			if health > 0 {
				// health > 0 means that we need at least that amount of health to come from this room to the exit
				dp[r][c] = health
			} else {
				// health <= 0 means that we were cured on the current cell more than we need,
				// so here it's enough to have 1hp
				dp[r][c] = 1
			}
		}
	}
	return dp[0][0]
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}
