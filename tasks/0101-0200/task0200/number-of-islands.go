package task0200

// Given an m x n 2D binary grid grid which represents a map of '1's (land) and '0's (water),
// return the number of islands.
// An island is surrounded by water and is formed by connecting adjacent lands horizontally
// or vertically. You may assume all four edges of the grid are all surrounded by water.
//
// Constraints:
//   m == grid.length
//   n == grid[i].length
//   1 <= m, n <= 300
//   grid[i][j] is '0' or '1'.
//
// Results:
//   Runtime: 2 ms, faster than 92.23% of Go online submissions for Number of Islands.
//   Memory Usage: 4.1 MB, less than 55.09% of Go online submissions for Number of Islands.
func numIslands(grid [][]byte) int {
	m, n := len(grid), len(grid[0])
	visited := make([][]bool, m)
	countOnes := 0
	for r := 0; r < m; r++ {
		visited[r] = make([]bool, n)
		for c := 0; c < n; c++ {
			if grid[r][c] == '1' {
				countOnes++
			}
		}
	}
	dirs := []Point{
		{0, 1},
		{0, -1},
		{1, 0},
		{-1, 0},
	}
	numIslands := 0
	stack := []Point{}
	var point Point
	var dir Point
	var x int
	var y int
	var nextX int
	var nextY int
	for x = 0; x < n; x++ {
		for y = 0; y < m; y++ {
			if grid[y][x] == '0' || visited[y][x] {
				continue
			}
			numIslands++
			stack = append(stack, Point{x, y})
			for len(stack) > 0 {
				point = stack[len(stack)-1]
				stack = stack[:len(stack)-1]
				visited[point.y][point.x] = true
				countOnes--
				if countOnes == 0 {
					return numIslands
				}
				for _, dir = range dirs {
					nextX, nextY = point.x+dir.x, point.y+dir.y

					if nextY >= 0 && nextY < m &&
						nextX >= 0 && nextX < n &&
						!visited[nextY][nextX] &&
						grid[nextY][nextX] == '1' {

						stack = append(stack, Point{nextX, nextY})
					}
				}
			}
		}
	}

	return numIslands
}

type Point struct {
	x int
	y int
}
