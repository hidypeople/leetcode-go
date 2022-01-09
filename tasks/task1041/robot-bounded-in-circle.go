package task1041

// On an infinite plane, a robot initially stands at (0, 0) and faces north. The robot can receive one of three instructions:
//   "G": go straight 1 unit;
//   "L": turn 90 degrees to the left;
//   "R": turn 90 degrees to the right.
// The robot performs the instructions given in order, and repeats them forever.
// Return true if and only if there exists a circle in the plane such that the robot never leaves the circle.
//
// Constraints:
//   1 <= instructions.length <= 100
//   instructions[i] is 'G', 'L' or, 'R'.
//
// Results:
//   Runtime: 0 ms, faster than 100.00% of Go online submissions for Robot Bounded In Circle.
//   Memory Usage: 2.1 MB, less than 29.55% of Go online submissions for Robot Bounded In Circle.
func isRobotBounded(instructions string) bool {
	// The idea is simple: we start from the point (0,0) with direction "east" and move according to the program.
	// We need to make 1, 2 or 4 program iterations to come back to the initial point.
	// We have 2 possible winning conditions:
	//   1. Come back to the (0,0) after completing the 1st programm run
	//   2. If we aren't at (0,0), we need to check the direction:
	//     - west        => loop has 2 robot program iterations
	//     - south/north => loop has 4 robot program iterations
	//     - east        => this is endless journey to the east and not the solution
	x, y := 0, 0
	directions := [][]int{{1, 0}, {0, -1}, {-1, 0}, {0, 1}}
	direction := 0
	for _, instruction := range instructions {
		switch instruction {
		case 'G':
			x += directions[direction][0]
			y += directions[direction][1]
		case 'L':
			direction = (direction + 1) % 4
		case 'R':
			direction = (direction + 3) % 4
		}
	}
	return direction > 0 || (x == 0 && y == 0)
}
