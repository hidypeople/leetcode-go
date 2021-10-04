package task0463

// You are given row x col grid representing a map where grid[i][j] = 1 represents land and grid[i][j] = 0 represents water.
// Grid cells are connected horizontally/vertically (not diagonally). The grid is completely surrounded by water,
// and there is exactly one island (i.e., one or more connected land cells).
// The island doesn't have "lakes", meaning the water inside isn't connected to the water around the island.
// One cell is a square with side length 1. The grid is rectangular, width and height don't exceed 100.
// Determine the perimeter of the island.
//
// Constraints:
//   row == grid.length
//   col == grid[i].length
//   1 <= row, col <= 100
//   grid[i][j] is 0 or 1.
//   There is exactly one island in grid.
//
// Results:
//   Runtime: 56 ms, faster than 81.13% of Go online submissions for Island Perimeter.
//   Memory Usage: 7.3 MB, less than 44.34% of Go online submissions for Island Perimeter.
func islandPerimeter(grid [][]int) int {
	m, n, result := len(grid), len(grid[0]), 0
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == 1 {
				// We found first '1': start from it
				stack := []int{i<<8 + j}
				for len(stack) > 0 {
					currentPoint := stack[len(stack)-1]
					r, c := currentPoint>>8, currentPoint%256
					stack = stack[:len(stack)-1]
					if grid[r][c] < 0 {
						continue
					}
					grid[r][c] = -1

					result += 4
					if r > 0 {
						// top
						if grid[r-1][c] > 0 {
							stack = append(stack, ((r-1)<<8)+c)
							result -= 1
						} else if grid[r-1][c] == -1 {
							result -= 1
						}
					}
					if c > 0 {
						// left
						if grid[r][c-1] > 0 {
							stack = append(stack, (r<<8)+c-1)
							result -= 1
						} else if grid[r][c-1] == -1 {
							result -= 1
						}
					}
					if r < m-1 {
						// bottom
						if grid[r+1][c] > 0 {
							stack = append(stack, ((r+1)<<8)+c)
							result -= 1
						} else if grid[r+1][c] == -1 {
							result -= 1
						}
					}
					if c < n-1 {
						if grid[r][c+1] > 0 {
							stack = append(stack, (r<<8)+c+1)
							result -= 1
						} else if grid[r][c+1] == -1 {
							result -= 1
						}
					}
				}
				return result
			}
		}
	}
	return 0
}
